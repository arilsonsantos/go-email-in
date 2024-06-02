package controller

import (
	"emailn/internal/internalerrors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	id := chi.URLParam(req, "id")
	idInt, _ := strconv.Atoi(id)
	campaign, err := h.CampaignService.GetBy(idInt)

	if campaign == nil {
		return nil, http.StatusNoContent, internalerrors.ErrNoContent
	}

	return campaign, http.StatusOK, err
}
