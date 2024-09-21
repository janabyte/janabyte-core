package handler

import (
	"net/http"
)

type Handler struct {
}

func NewUserHandler() *Handler {
    return &Handler{}
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {

}
