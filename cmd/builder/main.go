package main

import (
	"fmt"
	"time"

	"github.com/losev-al/GoLessons/pkg/httpclient"
)

func main() {
	client := httpclient.New().
		Host("https://nominatim.openstreetmap.org/").
		Path("?addressdetails=1&q=улица Ленинская Слобода Куликовский&format=json&limit=1").
		TimeOut(5 * time.Second).
		Build()
	fmt.Println(client.GetBody())
}
