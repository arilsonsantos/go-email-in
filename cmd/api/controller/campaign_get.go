package controller

import (
	"emailn/cmd/api/controller/utils"
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"github.com/go-chi/render"
	"net/http"
)

func CampaignsGet(service campaign.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var campaignDto contract.NewCampaignDto
		campaignDto, err := utils.ParseJSONRequest(req, campaignDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := service.CreateCampaign(campaignDto)
		if err != nil {
			utils.HandleError(w, req, err)
			return
		}

		render.Status(req, http.StatusCreated)
		render.JSON(w, req, map[string]string{"id": id})
	}
}
