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
	ctx context.Context
	DB  *sqlx.DB
}

func NewCampaignRepository(ctx context.Context, db *sqlx.DB) *CampaignRepository {
	return &CampaignRepository{
		ctx: ctx,
		DB:  db,
	}
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := c.DB.Queryx(queries.SELECT_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	campaignsMap := make(map[int]campaign.Campaign)
	for rows.Next() {
		var tempCampaign campaign.Campaign
		var tempContact campaign.Contact
		err = rows.Scan(&tempCampaign.ID, &tempCampaign.Name, &tempCampaign.CreatedAt, &tempCampaign.Content, &tempCampaign.Status, &tempContact.ID, &tempContact.Email)
		if err != nil {
			return nil, errors.New("erro ao executar a consulta")
		}
		currentCampaign, exists := campaignsMap[tempCampaign.ID]
		if exists {
			currentCampaign.Contacts = append(currentCampaign.Contacts, tempContact)
			campaignsMap[tempCampaign.ID] = currentCampaign
		} else {
			tempCampaign.Contacts = append(tempCampaign.Contacts, tempContact)
			campaignsMap[tempCampaign.ID] = tempCampaign
		}
	}

	campaignsToSliece := make([]campaign.Campaign, 0, len(campaignsMap))
	for _, c := range campaignsMap {
		campaignsToSliece = append(campaignsToSliece, c)
	}

	if err != nil {
		return nil, errors.New("erro ao executar a consulta")
	}

	return campaignsToSliece, nil
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

func (c *CampaignRepository) Save(ctx context.Context, campaign *campaign.Campaign) (int, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
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
