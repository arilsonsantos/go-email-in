package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
)

type Service struct {
	repository Repository
}

func (s *Service) CreateCampaign(dto contract.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(dto.Name, dto.Content, dto.Emails)
	if err != nil {
		return "", err
	}
	err = s.repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}
	return campaign.ID, nil
}
