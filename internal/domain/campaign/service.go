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
	campaign, err := NewCampaign(dto.Name, dto.Content, dto.Emails, dto.CreatedBy)
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
		campaignDto := contract.NewGetCampaignDto{
			ID:       campaign.ID,
			Name:     campaign.Name,
			Content:  campaign.Content,
			Contacts: s.getContactDtos(campaign.Contacts),
		}
		campaignDtos[i] = campaignDto
	}
	return &campaignDtos, nil
}

func (s *ServiceImpl) getContactDtos(contacts []Contact) []contract.NewGetContactDto {
	contactDtos := make([]contract.NewGetContactDto, len(contacts))
	for i, contact := range contacts {
		contactDtos[i] = contract.NewGetContactDto{
			Id:    contact.ID,
			Email: contact.Email,
		}
	}
	return contactDtos
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
