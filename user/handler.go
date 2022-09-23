package user

import (
	"encoding/json"
	"net/http"

	"github.com/ivanauliaa/learn-wire/domain"
)

type handler struct {
	svc domain.UserServiceInterface
}

func (h *handler) FetchByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")

			username := r.URL.Query()["username"][0]
			result, err := h.svc.FetchByUsername(r.Context(), username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			resultJson, _ := json.Marshal(result)
			w.Write(resultJson)

			return
		}

		http.Error(w, "Not found", http.StatusNotFound)
	}
}
