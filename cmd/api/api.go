package api

import (
	"context"
	"database/sql"
	"emailn/internal/controller"
	"emailn/internal/controller/utils"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Api() {

	r := chi.NewRouter()
	r.Use(middleware.Logger) // <--<< Logger should come before Recoverer
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	conn, _ := db.OpenConn()
	defer func(DB *sql.DB) { _ = DB.Close() }(conn.DB)

	ctx := context.Background()
	repo := database.NewCampaignRepository(ctx, conn.DB)

	service := campaign.ServiceImpl{
		Repository: repo,
	}

	handlers := controller.Handler{
		CampaignService: &service,
	}

	r.Post("/campaigns", controller.HandleError(handlers.CampaignPost))
	r.Get("/campaigns", controller.HandleError(handlers.CampaignGet))
	r.Get("/campaigns/{id}", controller.HandleError(handlers.CampaignGetById))

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		utils.HandleError500(err)
	}

}
