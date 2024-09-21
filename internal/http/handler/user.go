package handler

import (
	"net/http"

	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/service"
	"github.com/aidosgal/janabyte/janabyte-core/internal/utils"
)

type Handler struct {
    service service.UserService
}

func NewUserHandler(service service.UserService) *Handler {
    return &Handler{service}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleStore(w http.ResponseWriter, r *http.Request) {
    var user model.User
    if err := utils.ParseJSON(r, &user); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }
    
    err := h.service.CreateUser(user)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    utils.WriteJSON(w, http.StatusCreated, nil)
}
