package controller

import (
	"emailn/internal/internalerrors"
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Repository.Get()
	if campaigns == nil {
		return nil, http.StatusNoContent, internalerrors.ErrNoContent
	}
	return campaigns, http.StatusOK, err
}
