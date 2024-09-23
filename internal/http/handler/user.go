package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"

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

func (h *Handler) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleGetAllUsers"
	users, err := h.service.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	for _, user := range users {
		utils.WriteJSON(w, http.StatusOK, user)
	}

}

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleCreateUser"
	var user *model.User
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.CreateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	res := fmt.Sprintf("user with id: %d created", id)
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"result": res})
}

func (h *Handler) HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleDeleteUserById"
	IdParam := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (h *Handler) HandleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleUpdateUserById"
	IdParam := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(IdParam)
	if err != nil {
		log.Printf("error with url param")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var user *model.User
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.UpdateUserById(userId, user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleGetUserById"
	IdParam := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user, err := h.service.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
