package main

import (
	"fmt"
	"net/url"

	"github.com/hono3bono3/go-plactice/models"
)

func main() {
	var status models.HTTPStatus = 200
	fmt.Println(status.String())

	var route models.NationalRoute = 200
	fmt.Println(route.String())

	vs := url.Values{}
	vs.Add("hoge", "123")
	vs.Add("fuga", "456")
	vs.Add("fuga", "789")
	for k, v := range vs {
		fmt.Printf("%s: %+v\n", k, v)
	}

	fmt.Println(vs.Encode())

	// var consumers models.Consumers = []models.Consumer{{ID: "x01", IsActive: true}}
	consumers := models.Consumers{{ID: "x01", IsActive: true}, {ID: "x02", IsActive: false}, {ID: "x03", IsActive: true}}
	for _, v := range consumers.ActiveConsumer() {
		fmt.Printf("%+v", v)
	}
}
