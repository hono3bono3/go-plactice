package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Comment struct {
	Message  string
	UserName string
}

func Comments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments := make([]Comment, 0, 100)

	switch r.Method {
	case http.MethodGet:
		err := json.NewEncoder(w).Encode(comments)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
			return
		}

	}
}
