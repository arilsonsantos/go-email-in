package controller

import (
	"emailn/internal/contract"
	"emailn/internal/controller/utils"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	var campaignDto contract.NewCampaignDto
	_, _ = utils.ParseJSONRequest(req, &campaignDto)
	id, err := h.CampaignService.CreateCampaign(campaignDto)
	return map[string]string{"id": id}, http.StatusCreated, err
}
