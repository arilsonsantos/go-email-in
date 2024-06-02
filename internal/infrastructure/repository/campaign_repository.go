package repository

import (
	"context"
	"database/sql"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/queries"
	"emailn/internal/internalerrors"
	"errors"
	"fmt"
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

func (c *CampaignRepository) Get() (*[]campaign.Campaign, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rows, err := c.DB.Query(queries.SELECT_ALL)
	campaignResponsePtr, err := getCampaign(rows, err)

	return campaignResponsePtr, nil
}

func (c *CampaignRepository) GetBy(id int) (*campaign.Campaign, error) {
	rows, err := c.DB.Query(queries.SELECT_BY_ID, id)
	if err != nil {
		return nil, err
	}

	campaignResponsePtr, err := getCampaign(rows, err)

	if err != nil {
		return nil, err
	}

	campaignResponse := *campaignResponsePtr

	return &campaignResponse[0], nil
}

func (c *CampaignRepository) Save(ctx context.Context, campaingInput *campaign.Campaign) (int, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	tx, err := c.DB.Begin()
	var campaignId int
	err = c.DB.QueryRow(queries.INSERT_CAMPAIGN,
		campaingInput.Name,
		campaingInput.CreatedAt,
		campaingInput.Content,
		campaingInput.Status,
		campaingInput.CreatedBy).Scan(&campaignId)

	for _, contact := range campaingInput.Contacts {
		_, err = c.DB.Exec(queries.INSERT_CONTACT, contact.Email, campaignId)
	}
	err = tx.Commit()

	if err != nil {
		_ = tx.Rollback()
		fmt.Println("Error inserting campaign:", err)
		return 0, err
	}

	return campaignId, nil
}

func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	_, err := c.DB.Exec(queries.UPDATE_STATUS_CAMPAIGN, campaign.Status, campaign.ID)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func getCampaign(rows *sql.Rows, err error) (*[]campaign.Campaign, error) {
	campaignsMap := make(map[int]campaign.Campaign)
	for rows.Next() {
		var tempCampaign campaign.Campaign
		var tempContact campaign.Contact
		err = rows.Scan(
			&tempCampaign.ID,
			&tempCampaign.Name,
			&tempCampaign.CreatedAt,
			&tempCampaign.Content,
			&tempCampaign.Status,
			&tempContact.ID,
			&tempContact.Email)

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

	return &campaignsToSlice, nil
}
