package controller

import (
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, req *http.Request) {
	render.Status(req, http.StatusOK)
	render.JSON(w, req, h.CampaignService.Repository.Get())
}
