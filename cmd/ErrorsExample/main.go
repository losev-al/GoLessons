package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const minWait = 10000
const avgWait = 100

var log io.Writer = os.Stdout

type bodyLength struct {
	length int
	err    error
}

func getBodyLength(ctx context.Context, url string) (int, error) {
	ch := make(chan bodyLength)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- bodyLength{length: 0, err: err}
	}
	tr := &http.Transport{}
	client := &http.Client{Transport: tr, Timeout: time.Duration(rand.Intn(avgWait)+minWait) * time.Millisecond}
	go func(ch chan bodyLength) {
		resp, err := client.Do(req)
		if err != nil {
			ch <- bodyLength{length: 0, err: err}
			return
		}
		defer resp.Body.Close()
		bodyData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ch <- bodyLength{length: 0, err: err}
			return
		}
		ch <- bodyLength{length: len(bodyData), err: err}
	}(ch)
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		return 0, errors.New("Cancel by context")
	case res := <-ch:
		return res.length, res.err
	}
}

func logTimings(path string, start time.Time, out io.Writer) {
	out.Write([]byte(fmt.Sprintf("path: %v time: %v\n", path, time.Since(start))))
}

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx, _ = context.WithTimeout(ctx, time.Duration(rand.Intn(avgWait)+minWait)*time.Millisecond)
	res, err := getBodyLength(ctx, "http://ya.ru")
	if err != nil {
		log.Write([]byte(fmt.Sprintf("full err: %+v\n", err)))
		switch err := errors.Cause(err).(type) {
		case *url.Error:
			log.Write([]byte(fmt.Sprintf("resource %s err: %+v\n", err.URL, err.Err)))
			http.Error(w, "remote resource error", 500)
		default:
			fmt.Printf("%+v\n", err)
			http.Error(w, "parsing error", 500)
		}
		return
	}
	w.Write([]byte(strconv.FormatInt(int64(res), 10)))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/", handler)

	siteHandler := timingMiddleware(siteMux)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", siteHandler)
}

func timingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer logTimings(r.URL.Path, time.Now(), log)
		next.ServeHTTP(w, r)
	})
}
