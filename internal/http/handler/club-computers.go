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

type ComputerHandler struct {
	service service.ComputerService
}

func NewComputerHandler(service service.ComputerService) *ComputerHandler {
	return &ComputerHandler{service: service}
}

func (h *ComputerHandler) HandlerCreateComputer(w http.ResponseWriter, r *http.Request) {
	var computer *model.Computers
	if err := utils.ParseJSON(r, &computer); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.CreateComputer(computer)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := fmt.Sprintf("Computer with id:%d created", id)
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"res": res})
}

func (h *ComputerHandler) HandlerGetComputerById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	computer, err := h.service.GetComputerById(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, computer)
}

func (h *ComputerHandler) HandlerGetAllComputers(w http.ResponseWriter, r *http.Request) {
	computers, err := h.service.GetAllComputers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	for _, computer := range computers {
		utils.WriteJSON(w, http.StatusOK, computer)
	}
}

func (h *ComputerHandler) HandlerDeleteComputerById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteComputerById(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"res": "computer successfully deleted"})
}

func (h *ComputerHandler) HandlerUpdateComputerById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var computer *model.Computers
	if err := utils.ParseJSON(r, &computer); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.UpdateComputer(id, computer)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"res": "computer successfully updated"})
}
