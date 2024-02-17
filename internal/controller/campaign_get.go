package controller

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Repository.Get()
	return campaigns, http.StatusOK, err
}
