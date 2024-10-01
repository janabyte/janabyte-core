package handler

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/service"
	"github.com/aidosgal/janabyte/janabyte-core/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type HandlerRole struct {
	service service.RoleService
}

func NewHandlerRole(service service.RoleService) *HandlerRole {
	return &HandlerRole{service: service}
}

func (h *HandlerRole) HandlerCreateRole(w http.ResponseWriter, r *http.Request) {
	var role *model.Roles
	if err := utils.ParseJSON(r, &role); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.CreateRole(role)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := fmt.Sprintf("User with id: %d created", id)
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"res": res})

}

func (h *HandlerRole) HandlerDeleteRole(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	roleId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteRole(roleId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"res": "successfully deleted"})
}

func (h *HandlerRole) HandlerGetRoleById(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	roleId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	role, err := h.service.GetRoleById(roleId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, role)
}

func (h *HandlerRole) HandlerGetAllRole(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllRoles()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	for _, user := range users {
		utils.WriteJSON(w, http.StatusOK, user)
	}

}
func (h *HandlerRole) HandlerUpdateRoleById(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	roleId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var role *model.Roles
	if err := utils.ParseJSON(r, &role); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.UpdateRoleById(roleId, role)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"res": "successfully updated"})
}
