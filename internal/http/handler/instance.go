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

type InstanceHandler struct {
	service service.InstanceService
}

func NewInstanceHandler(service service.InstanceService) *InstanceHandler {
	return &InstanceHandler{service: service}
}

func (h *InstanceHandler) HandlerCreateInstance(w http.ResponseWriter, r *http.Request) {
	var instance *model.Instance
	if err := utils.ParseJSON(r, &instance); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.CreateInstance(instance)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := fmt.Sprintf("instance with id: %d successfully created", id)
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"res": res})

}

func (h *InstanceHandler) HandlerGetInstanceById(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	clubId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	instance, err := h.service.GetInstanceById(clubId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, instance)
}

func (h *InstanceHandler) HandlerGetAllInstance(w http.ResponseWriter, r *http.Request) {
	instances, err := h.service.GetAllInstances()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	for _, instance := range instances {
		utils.WriteJSON(w, http.StatusOK, instance)
	}
}

func (h *InstanceHandler) HandlerDeleteInstanceById(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	clubId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteInstanceById(clubId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"res": "successfully deleted"})
}

func (h *InstanceHandler) HandlerUpdateInstanceById(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	clubId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var instance *model.Instance
	if err = utils.ParseJSON(r, &instance); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.UpdateInstanceById(clubId, instance)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"res": "successfully updated"})
}
