package database

import (
	"context"
	"database/sql"
	"emailn/internal/domain/campaign"
	"errors"
	"time"
)

type CampaignRepository struct {
	DB *sql.DB
}

func NewCampaignRepository(db *sql.DB) *CampaignRepository {
	return &CampaignRepository{DB: db}
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := "SELECT id, name FROM campaign"
	rows, err := c.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, errors.New("erro ao executar a consulta")
	}
	defer rows.Close()

	var campaigns []campaign.Campaign
	for rows.Next() {
		var campaign2 campaign.Campaign
		err := rows.Scan(&campaign2.ID, &campaign2.Name) // Adicione outras colunas conforme necessário
		if err != nil {
			return nil, errors.New("erro ao escanear a linha")
		}
		campaigns = append(campaigns, campaign2)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("erro ao percorrer as linhas do resultado")
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
