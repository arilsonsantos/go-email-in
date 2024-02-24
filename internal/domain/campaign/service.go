package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
)

type Service interface {
	CreateCampaign(dto contract.NewCampaignDto) (string, error)
	GetCampaigns() ([]Campaign, error)
	GetBy(id string) (*contract.NewCampaignResponseDto, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) CreateCampaign(dto contract.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(dto.Name, dto.Content, dto.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}
	return campaign.ID, nil
}

func (s *ServiceImpl) GetCampaigns() ([]Campaign, error) {
	return s.Repository.Get()
}

func (s *ServiceImpl) GetBy(id string) (*contract.NewCampaignResponseDto, error) {
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
