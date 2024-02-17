package controller

import (
	"emailn/internal/controller/utils"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, req *http.Request) {
	campaigns, err := h.CampaignService.Repository.Get()
	if err != nil {
		utils.HandleError(err)

	}
	render.Status(req, http.StatusOK)
	render.JSON(w, req, campaigns)
}
