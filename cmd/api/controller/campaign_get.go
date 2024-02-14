package controller

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/internalerrors"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

func CampaignsGet(service campaign.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		request, err := parseJSONRequest(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := service.CreateCampaign(request)
		if err != nil {
			handleError(w, req, err)
			return
		}

		render.Status(req, http.StatusCreated)
		render.JSON(w, req, map[string]string{"id": id})
	}
}

func parseJSONRequest(req *http.Request) (contract.NewCampaignDto, error) {
	var request contract.NewCampaignDto
	err := render.DecodeJSON(req.Body, &request)
	return request, err
}

func handleError(w http.ResponseWriter, req *http.Request, err error) {
	if errors.Is(err, internalerrors.ErrInternal) {
		render.Status(req, http.StatusInternalServerError)
	} else {
		render.Status(req, http.StatusBadRequest)
	}
	render.JSON(w, req, map[string]string{"error": err.Error()})
}
