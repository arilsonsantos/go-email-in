package campaign

import "emailn/internal/contract"

type Repository interface {
	Save(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetBy(id string) (*contract.NewCampaignResponseDto, error)
}
