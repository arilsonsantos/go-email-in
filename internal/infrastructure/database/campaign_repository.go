package database

import (
	"context"
	"database/sql"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/queries"
	"errors"
	"fmt"
	"log"
	"time"
)

type CampaignRepository struct {
	ctx context.Context
	DB  *sql.DB
}

func NewCampaignRepository(ctx context.Context, db *sql.DB) *CampaignRepository {
	return &CampaignRepository{
		ctx: ctx,
		DB:  db,
	}
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := c.DB.Query(queries.SELECT_ALL)

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

	campaignsToSlice := make([]campaign.Campaign, 0, len(campaignsMap))
	for _, c := range campaignsMap {
		campaignsToSlice = append(campaignsToSlice, c)
	}

	if err != nil {
		return nil, errors.New("erro ao executar a consulta")
	}

	return campaignsToSlice, nil
}

func (c *CampaignRepository) GetBy(id int) (*campaign.Campaign, error) {
	var campaignResponse campaign.Campaign
	err := c.DB.QueryRow(queries.SELECT_BY_ID, id).Scan(
		&campaignResponse.ID,
		&campaignResponse.Name,
		&campaignResponse.CreatedAt,
		&campaignResponse.Content,
		&campaignResponse.Status)

	if err != nil {
		return nil, err
	}

	return &campaignResponse, nil
}

func (c *CampaignRepository) Save(ctx context.Context, campaign *campaign.Campaign) (int, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	//params := map[string]interface{}{"name": campaign.Name}

	result, err := c.DB.ExecContext(ctx, queries.INSERT_CAMPAIGN_NAME, campaign)

	if err != nil {
		fmt.Println("Error inserting campaign:", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	log.Println("Rows affected:", rowsAffected)

	var x, _ = result.LastInsertId()
	return int(x), nil
}
