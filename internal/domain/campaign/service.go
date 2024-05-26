package campaign

import (
	"context"
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
)

type Service interface {
	CreateCampaign(ctx context.Context, dto contract.NewCampaignDto) (int, error)
	GetCampaigns() (*[]contract.NewCampaignResponseDto, error)
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

	result, err := s.Repository.Save(ctx, campaign)
	if err != nil {
		return 0, internalerrors.ErrInternal
	}
	return result, nil
}

func (s *ServiceImpl) GetCampaigns() (*[]contract.NewCampaignResponseDto, error) {
	campaigns, _ := s.Repository.Get()
	campaignDtos := make([]contract.NewCampaignResponseDto, len(*campaigns))

	for i, campaign := range *campaigns {
		var campaignDto contract.NewCampaignResponseDto
		contactDtos := make([]contract.NewContactDto, len(campaign.Contacts))

		for i, contact := range campaign.Contacts {
			var contactDto contract.NewContactDto
			contactDto.Id = contact.ID
			contactDto.Email = contact.Email
			contactDtos[i] = contactDto
		}

		campaignDto.ID = campaign.ID
		campaignDto.Name = campaign.Name
		campaignDto.Content = campaign.Content
		campaignDto.Contacts = contactDtos

		campaignDtos[i] = campaignDto
	}

	return &campaignDtos, nil
}

func (s *ServiceImpl) GetBy(id int) (*contract.NewCampaignResponseDto, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	contacts := make([]contract.NewContactDto, len(campaign.Contacts))
	for i, contact := range campaign.Contacts {
		contacts[i] = contract.NewContactDto{
			Id:    contact.ID,
			Email: contact.Email,
		}
	}
	return &contract.NewCampaignResponseDto{
		ID:       campaign.ID,
		Name:     campaign.Name,
		Content:  campaign.Content,
		Status:   campaign.Status,
		Contacts: contacts,
	}, nil
}
