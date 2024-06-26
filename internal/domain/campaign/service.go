package campaign

import (
	"context"
	contract2 "emailn/internal/controller/dto"
	"emailn/internal/internalerrors"
	"errors"
)

type Service interface {
	CreateCampaign(ctx context.Context, dto contract2.NewPostCampaignDto) (int, error)
	GetCampaigns() (*[]contract2.NewGetCampaignDto, error)
	GetBy(id int) (*contract2.NewGetCampaignDto, error)
	Start(id int) error
}

type ServiceImpl struct {
	Repository Repository
	SendEmail  func(campaign *Campaign) error
}

func (s *ServiceImpl) CreateCampaign(ctx context.Context, dto contract2.NewPostCampaignDto) (int, error) {
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

func (s *ServiceImpl) GetCampaigns() (*[]contract2.NewGetCampaignDto, error) {
	campaigns, _ := s.Repository.Get()
	campaignDtos := make([]contract2.NewGetCampaignDto, len(*campaigns))

	for i, campaign := range *campaigns {
		campaignDto := getCampaign(&campaign, s)
		campaignDtos[i] = campaignDto
	}
	return &campaignDtos, nil
}

func (s *ServiceImpl) GetBy(id int) (*contract2.NewGetCampaignDto, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	var campaignDto = getCampaign(campaign, s)

	return &campaignDto, nil
}

func (s *ServiceImpl) Start(id int) error {
	campaign, err := getValidCampaign(id, s)
	if err != nil {
		return err
	}

	go sendEmail(err, s, campaign)

	campaign.Started()
	err = s.Repository.Update(campaign)
	return nil
}

func sendEmail(err error, s *ServiceImpl, campaign *Campaign) {
	err = s.SendEmail(campaign)
	if err != nil {
		campaign.Failed()
	} else {
		campaign.Done()
	}
	err = s.Repository.Update(campaign)
}

func getValidCampaign(id int, s *ServiceImpl) (*Campaign, error) {
	campaign, err := s.Repository.GetBy(id)
	if err != nil {
		return nil, internalerrors.ErrNotFound
	}

	if campaign.Status != Pending {
		return nil, errors.New("invalid status")
	}
	return campaign, err
}

func getCampaign(campaign *Campaign, s *ServiceImpl) contract2.NewGetCampaignDto {
	return contract2.NewGetCampaignDto{
		ID:       campaign.ID,
		Name:     campaign.Name,
		Content:  campaign.Content,
		Status:   campaign.Status,
		Contacts: s.getContactDtos(campaign.Contacts),
	}
}

func (s *ServiceImpl) getContactDtos(contacts []Contact) []contract2.NewGetContactDto {
	contactDtos := make([]contract2.NewGetContactDto, len(contacts))
	for i, contact := range contacts {
		contactDtos[i] = contract2.NewGetContactDto{
			Id:    contact.ID,
			Email: contact.Email,
		}
	}
	return contactDtos
}
