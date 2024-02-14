package api

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/internalerrors"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
)

func Api() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	r.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaignDto
		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := service.CreateCampaign(request)

		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
			} else {
				render.Status(r, http.StatusBadRequest)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]string{"id": id})

	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}

}
