package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *Service
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newU UserMini
	json.NewDecoder(r.Body).Decode(&newU)

	code := h.Service.AppendUser(newU)

	w.WriteHeader(code)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser UserMini
	json.NewDecoder(r.Body).Decode(&loginUser)

	code, jwt := h.Service.LoginUser(loginUser)

	w.WriteHeader(code)
	if code == http.StatusOK {
		w.Write([]byte(jwt))
	}
}
