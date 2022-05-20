package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/hono3bono3/go-plactice/server/handler"
	"github.com/hono3bono3/go-plactice/server/store"
)

func Serve() {
	var mutex = &sync.RWMutex{}

	comments := store.GetComments()

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			mutex.RLock()
			err := json.NewEncoder(w).Encode(comments)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()

		case http.MethodPost:
			var c handler.Comment
			err := json.NewDecoder(r.Body).Decode(&c)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.Lock()
			comments = append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status": "created"}`))
		default:
			http.Error(w, `{"status": "permits only GET or POST"}`, http.StatusMethodNotAllowed)
		}

	})

	http.HandleFunc("/hello", handler.Hello)

	fmt.Println("Start listening at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "")
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}
