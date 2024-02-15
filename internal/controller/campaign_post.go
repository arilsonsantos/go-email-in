package controller

import (
	"emailn/internal/contract"
	"emailn/internal/controller/utils"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, req *http.Request) {
	var campaignDto contract.NewCampaignDto
	campaignDto, err := utils.ParseJSONRequest(req, campaignDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := h.CampaignService.CreateCampaign(campaignDto)
	if err != nil {
		utils.HandleError(w, req, err)
		return
	}

	render.Status(req, http.StatusCreated)
	render.JSON(w, req, map[string]string{"id": id})
}
