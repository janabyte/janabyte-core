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

type HandlerClub struct {
	service service.ClubService
}

func NewClubHandler(service service.ClubService) *HandlerClub {
	return &HandlerClub{service: service}

}

func (h *HandlerClub) HandlerCreateClub(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandlerCreateClub"
	var club *model.Club
	if err := utils.ParseJSON(r, &club); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.CreateClub(club)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := fmt.Sprintf("club with id: %d created", id)
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"result": res})

}

func (h *HandlerClub) HandlerGetClubById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandlerGetClubById"
	IdParam := chi.URLParam(r, "id")
	clubId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	club, err := h.service.GetClubById(clubId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, club)
}
func (h *HandlerClub) HandlerGetClubList(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandlerGetClubList"
	clubs, err := h.service.GetAllClubs()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	for _, v := range clubs {
		utils.WriteJSON(w, http.StatusOK, v)
	}

}

func (h *HandlerClub) HandlerUpdateClub(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandlerUpdateClub"
	var club *model.Club
	IdParam := chi.URLParam(r, "id")
	clubId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.ParseJSON(r, &club); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.UpdateClub(clubId, club)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, club)
}
func (h *HandlerClub) HandlerDeleteClub(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandlerDeleteClub"
	IdParam := chi.URLParam(r, "id")
	clubId, err := strconv.Atoi(IdParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteClub(clubId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "successfully deleted"})
}
