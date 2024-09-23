package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
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

}

func CreateUser(creator service.UserManipulator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := utils.ParseJSON(r, &user); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		err := creator.CreateUser(&user)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, user)
	}
}

func GetAllUsers(getter service.UserManipulator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetAllUsers"
		users, err := getter.GetAllUsers()
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		for _, user := range users {
			utils.WriteJSON(w, http.StatusOK, user)
		}
	}
}

func GetUserById(getter service.UserManipulator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetUserById"
		IdParam := chi.URLParam(r, "id")
		userId, err := strconv.Atoi(IdParam)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user id: %s %s", IdParam, op))
			return
		}
		user, err := getter.GetUserById(userId)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user %d not found %s", userId, op))
			return
		}
		utils.WriteJSON(w, http.StatusOK, user)
	}
}

func DeleteUserById(deleter service.UserManipulator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.DeleteUserById"
		IdParam := chi.URLParam(r, "id")
		userId, err := strconv.Atoi(IdParam)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user id: %s %s", IdParam, op))
			return
		}
		err = deleter.DeleteUser(userId)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user %d not found %s", userId, op))
			return
		}
		utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}

func UpdateUserById(updater service.UserManipulator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.UpdateUserById"
		//IdParam := chi.URLParam(r, "id")
		//, err := strconv.Atoi(IdParam)
		//if err != nil {
		//	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user id: %s %s", IdParam, op))
		//	return
		//}
		var user *model.User
		if err := utils.ParseJSON(r, &user); err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid json: %s %s", err, op))

			return
		}
		err := updater.UpdateUser(user)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user %d not found %s", user.Id, err))
			return
		}
		utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}

//func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
//
//}

//func (h *Handler) HandleStore(w http.ResponseWriter, r *http.Request) {
//	var user model.User
//	if err := utils.ParseJSON(r, &user); err != nil {
//		utils.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	err := h.service.CreateUser(user)
//	if err != nil {
//		utils.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	utils.WriteJSON(w, http.StatusCreated, nil)
//}
