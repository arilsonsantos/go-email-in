package campaign

import (
	"context"
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
)

type Service interface {
	CreateCampaign(ctx context.Context, dto contract.NewCampaignDto) (int, error)
	GetCampaigns() ([]Campaign, error)
	GetBy(id int) (*contract.NewCampaignResponseDto, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) CreateCampaign(ctx context.Context, dto contract.NewCampaignDto) (int, error) {
	campaign, err := NewCampaign(dto.Name, dto.Content, dto.Emails)
	if err != nil {
		return 0, err
	}
	var result int
	result, err = s.Repository.Save(ctx, campaign)
	if err != nil {
		return 0, internalerrors.ErrInternal
	}
	return result, nil
}

func (s *ServiceImpl) GetCampaigns() ([]Campaign, error) {
	return s.Repository.Get()
}

func (s *ServiceImpl) GetBy(id int) (*contract.NewCampaignResponseDto, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.NewCampaignResponseDto{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
