package main

import (
	"fmt"

	"github.com/hono3bono3/go-plactice/models"
)

func main() {
	var status models.HTTPStatus = 200
	fmt.Println(status.String())

	var route models.NationalRoute = 200
	fmt.Println(route.String())
}
