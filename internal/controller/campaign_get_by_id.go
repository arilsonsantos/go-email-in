package controller

import (
	"emailn/internal/internalerrors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	id := chi.URLParam(req, "id")

	campaign, err := h.CampaignService.GetBy(id)
	if campaign == nil {
		return nil, http.StatusNoContent, internalerrors.ErrNoContent
	}
	return campaign, http.StatusOK, err
}
