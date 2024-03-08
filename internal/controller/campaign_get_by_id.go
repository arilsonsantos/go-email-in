package controller

import (
	"emailn/internal/internalerrors"
	"net/http"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	id := req.URL.Query().Get("id")
	campaign, err := h.CampaignService.GetBy(id)
	if campaign == nil {
		return nil, http.StatusNoContent, internalerrors.ErrNoContent
	}
	return campaign, http.StatusOK, err
}
