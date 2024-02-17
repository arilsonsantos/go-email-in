package controller

import (
	"emailn/internal/contract"
	"emailn/internal/controller/utils"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	var campaignDto contract.NewCampaignDto
	campaignDto, err := utils.ParseJSONRequest(req, campaignDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, 500, err
	}

	id, err := h.CampaignService.CreateCampaign(campaignDto)
	if err != nil {
		var httpCode, err = utils.HandleError(err)
		return nil, httpCode, err
	}

	return map[string]string{"id": id}, http.StatusCreated, nil
}
