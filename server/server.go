package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hono3bono3/go-plactice/server/handler"
)

func Serve() {

	http.HandleFunc("/hello", handler.Hello)

	fmt.Println("Start listening at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "")
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}
