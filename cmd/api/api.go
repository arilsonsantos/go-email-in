package api

import (
	"emailn/cmd/api/controller"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.HandleFunc("/campaigns", controller.CampaignsGet(service))

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}

}
