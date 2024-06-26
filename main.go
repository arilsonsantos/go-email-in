package main

import (
	"context"
	"database/sql"
	"emailn/internal/controller"
	"emailn/internal/controller/utils"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/db"
	"emailn/internal/infrastructure/email"
	"emailn/internal/infrastructure/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger) // <--<< Logger should come before Recoverer
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	conn, _ := db.OpenConn()
	defer func(DB *sql.DB) { _ = DB.Close() }(conn.DB)

	ctx := context.Background()
	repo := repository.NewCampaignRepository(ctx, conn.DB)

	service := campaign.ServiceImpl{
		Repository: repo,
		SendEmail:  email.SendEmail,
	}

	handlers := controller.Handler{
		CampaignService: &service,
	}

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(controller.Auth)
		r.Post("/", controller.HandleError(handlers.CampaignPost))
	})

	r.Post("/campaigns/open", controller.HandleError(handlers.CampaignPost))

	r.Get("/campaigns", controller.HandleError(handlers.CampaignGet))
	r.Get("/campaigns/{id}", controller.HandleError(handlers.CampaignGetById))
	r.Patch("/campaigns/start/{id}", controller.HandleError(handlers.CampaignStart))

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		utils.HandleError500(err)
	}
}
