package httpclient

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/valyala/fasthttp"
)

// Client ...
type Client interface {
	GetBody() (body string, err error)
}

type client struct {
	host       string
	path       string
	logger     log.Logger
	timeOut    time.Duration
}

// GetBody ...
func (c *client) GetBody() (body string, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(fmt.Sprintf("%s%s", c.host, c.path))
	req.Header.SetMethod(fasthttp.MethodGet)

	client := &fasthttp.Client{}
	if c.timeOut == 0 {
		err = client.Do(req, resp)
	} else {
		err = client.DoTimeout(req, resp, c.timeOut)
	}
	if  err != nil {
		_ = c.logError("msg", "something went wrong with purchase request", "err", err)
		return
	}
	body = string(resp.Body())
	return
}

func (c *client) logError(keyvals ...interface{}) (err error) {
	if c.logger != nil {
		err = level.Error(c.logger).Log(keyvals...)
	}
	return
}
