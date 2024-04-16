package database

import (
	"context"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/queries"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type CampaignRepository struct {
	DB *sqlx.DB
}

func NewCampaignRepository(db *sqlx.DB) *CampaignRepository {
	return &CampaignRepository{DB: db}
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var campaigns []campaign.Campaign
	err := c.DB.SelectContext(ctx, &campaigns, queries.SELECT_ID_NAME)

	if err != nil {
		return nil, errors.New("erro ao executar a consulta")
	}

	return campaigns, nil
}

func (c *CampaignRepository) GetBy(id int) (*campaign.Campaign, error) {
	var campaignResponse campaign.Campaign
	//params := map[string]interface{}{"id": id}
	err := c.DB.Get(&campaignResponse, queries.SELECT_ID_NAME_BY_ID, id)

	if err != nil {
		return nil, err
	}

	return &campaignResponse, nil
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//params := map[string]interface{}{"name": campaign.Name}

	result, err := c.DB.NamedExecContext(ctx, queries.INSERT_CAMPAIGN_NAME, campaign)

	if err != nil {
		fmt.Println("Error inserting campaign:", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	log.Println("Rows affected:", rowsAffected)

	var x, _ = result.LastInsertId()
	return int(x), nil
}
