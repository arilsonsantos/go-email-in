package api

import (
	"emailn/internal/controller"
	"emailn/internal/controller/utils"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Api() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	service := campaign.ServiceImpl{
		Repository: &database.CampaignRepository{},
	}

	handlers := controller.Handler{
		CampaignService: &service,
	}

	r.Post("/campaigns", controller.HandleControllerError(handlers.CampaignPost))
	r.Get("/campaigns", controller.HandleControllerError(handlers.CampaignGet))
	r.Get("/campaign/{id}", controller.HandleControllerError(handlers.CampaignGetById))

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		utils.HandleError500(err)
	}

}
