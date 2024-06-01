package controller

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) CampaignStart(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	id := chi.URLParam(req, "id")
	idInt, _ := strconv.Atoi(id)
	err := h.CampaignService.Start(idInt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	return nil, http.StatusOK, err
}
