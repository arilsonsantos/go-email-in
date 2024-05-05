package api

import (
	"context"
	"emailn/internal/controller"
	"emailn/internal/controller/utils"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/db"
	"net/http"

	"github.com/jmoiron/sqlx"

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
	defer func(DB *sqlx.DB) { _ = DB.Close() }(conn.DB)

	//initDB(conn.DB)

	ctx := context.Background()
	repo := database.NewCampaignRepository(ctx, conn.DB)

	service := campaign.ServiceImpl{
		Repository: repo,
	}

	handlers := controller.Handler{
		CampaignService: &service,
	}

	//initDB(dbConn)

	r.Post("/campaigns", controller.HandleError(handlers.CampaignPost))
	r.Get("/campaigns", controller.HandleError(handlers.CampaignGet))
	r.Get("/campaigns/{id}", controller.HandleError(handlers.CampaignGetById))

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		utils.HandleError500(err)
	}

}

func initDB(dbConn *sqlx.DB) {
	_, _ = dbConn.Exec(`
		CREATE TABLE IF NOT EXISTS Contact (
	  	id TEXT PRIMARY KEY,
	  	email TEXT)
	`)

	_, _ = dbConn.Exec(`
		CREATE TABLE IF NOT EXISTS Campaign (
        id TEXT PRIMARY KEY,
        name TEXT,
        created_at DATETIME,
        content TEXT,
        contact_ID TEXT,
        status TEXT
        --FOREIGN KEY (ContactID) REFERENCES Contact(ID))   	
	`)
}
