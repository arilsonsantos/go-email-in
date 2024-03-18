package database

import (
	"context"
	"emailn/internal/domain/campaign"
	"errors"
	"github.com/jmoiron/sqlx"
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
	query := "SELECT id, name FROM campaign"
	var campaigns []campaign.Campaign
	err := c.DB.SelectContext(ctx, &campaigns, query)

	if err != nil {
		return nil, errors.New("erro ao executar a consulta")
	}

	return campaigns, nil
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	campaignResponse, err := campaign.NewCampaign("Nome A", "Conteúdo A", []string{"item1", "item2", "item3"})
	if err != nil {
		return nil, err

	}
	return campaignResponse, nil
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {

	return nil
}
