package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/hono3bono3/go-plactice/models"
	"github.com/hono3bono3/go-plactice/server/store"
)

func Comments(w http.ResponseWriter, r *http.Request) {
	mutex := &sync.RWMutex{}
	comments := store.GetComments()

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
		var c models.Comment
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
			return
		}
		mutex.Lock()
		store.AddComment(c)
		mutex.Unlock()

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status": "created"}`))
	default:
		http.Error(w, `{"status": "permits only GET or POST"}`, http.StatusMethodNotAllowed)
	}
}
