package handler

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/token"
	"github.com/aidosgal/janabyte/janabyte-core/internal/logger"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/service"
	"github.com/aidosgal/janabyte/janabyte-core/internal/utils"
)

const (
	access_duration  = time.Minute * 15
	refresh_duration = time.Hour * 24
)

var (
	slogger = logger.SetupLogger()
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

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleAuthenticateUserByLogin"

	var user *model.User
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	retUser, err := h.service.AuthenticateUser(r, user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if retUser == nil && err == nil {
		utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "successfully logged in"})
		return
	}
	//todo: use real permission model
	permission := []string{"ADMIN"} //пока так оставил, не понял пока как передавать permission юзерам
	//slogger.Debug("retUser: ", retUser)
	accessToken, accessClaims, err := token.MakeToken(retUser, permission, access_duration)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("cannot create access token"))
		return
	}
	refreshToken, refreshClaims, err := token.MakeToken(retUser, permission, refresh_duration)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("cannot create refresh token"))
		return
	}
	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Unix(accessClaims.ExpiresAt, 0),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, accessCookie)

	refreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Unix(refreshClaims.ExpiresAt, 0),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, refreshCookie)

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"result": "user successfully authenticated",
		"token":  accessToken,
	})

}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	const op = "handler.Logout"
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})
	utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "successfully logged out"})
}
func (h *Handler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	const op = "handler.RefreshToken"
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	refreshClaims, err := token.VerifyToken(refreshCookie.Value)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	//todo: use real permission model
	permission := []string{"ADMIN"} //пока так оставил, не понял пока как передавать permission юзерам

	user, err := h.service.GetUserById(refreshClaims.Id)
	accessToken, accessClaims, err := token.MakeToken(user, permission, access_duration)
	if err != nil {
		http.Error(w, "error when creating access token in login", 404)
		return
	}
	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Unix(accessClaims.ExpiresAt, 0),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, accessCookie)
	utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "Extended session"})
}
