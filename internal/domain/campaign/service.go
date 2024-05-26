package campaign

import (
	"context"
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
)

type Service interface {
	CreateCampaign(ctx context.Context, dto contract.NewPostCampaignDto) (int, error)
	GetCampaigns() (*[]contract.NewGetCampaignDto, error)
	GetBy(id int) (*contract.NewGetCampaignDto, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) CreateCampaign(ctx context.Context, dto contract.NewPostCampaignDto) (int, error) {
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

func (s *ServiceImpl) GetCampaigns() (*[]contract.NewGetCampaignDto, error) {
	campaigns, _ := s.Repository.Get()
	campaignDtos := make([]contract.NewGetCampaignDto, len(*campaigns))

	for i, campaign := range *campaigns {
		var campaignDto contract.NewGetCampaignDto
		contactDtos := make([]contract.NewGetContactDto, len(campaign.Contacts))

		for i, contact := range campaign.Contacts {
			var contactDto contract.NewGetContactDto
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

func (s *ServiceImpl) GetBy(id int) (*contract.NewGetCampaignDto, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	contacts := make([]contract.NewGetContactDto, len(campaign.Contacts))
	for i, contact := range campaign.Contacts {
		contacts[i] = contract.NewGetContactDto{
			Id:    contact.ID,
			Email: contact.Email,
		}
	}
	return &contract.NewGetCampaignDto{
		ID:       campaign.ID,
		Name:     campaign.Name,
		Content:  campaign.Content,
		Status:   campaign.Status,
		Contacts: contacts,
	}, nil
}
