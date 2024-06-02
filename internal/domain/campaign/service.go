package campaign

import (
	"context"
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
	"errors"
)

type Service interface {
	CreateCampaign(ctx context.Context, dto contract.NewPostCampaignDto) (int, error)
	GetCampaigns() (*[]contract.NewGetCampaignDto, error)
	GetBy(id int) (*contract.NewGetCampaignDto, error)
	Start(id int) error
}

type ServiceImpl struct {
	Repository Repository
	SendEmail  func(campaign *Campaign) error
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
		campaignDto := getCampaign(&campaign, s)
		campaignDtos[i] = campaignDto
	}
	return &campaignDtos, nil
}

func (s *ServiceImpl) GetBy(id int) (*contract.NewGetCampaignDto, error) {
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
	err = s.SendEmail(campaign)
	if err != nil {
		return internalerrors.ErrSendingEmail
	}

	campaign.Done()
	err = s.Repository.Update(Done, id)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
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

func getCampaign(campaign *Campaign, s *ServiceImpl) contract.NewGetCampaignDto {
	return contract.NewGetCampaignDto{
		ID:       campaign.ID,
		Name:     campaign.Name,
		Content:  campaign.Content,
		Status:   campaign.Status,
		Contacts: s.getContactDtos(campaign.Contacts),
	}
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
