package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	campaign = contract.NewCampaignDto{
		Name:    "My campaign",
		Content: "Body of the campaign",
		Emails: []string{
			"teste1@email.com",
		},
	}
	repository = new(repositoryMock)
	service    = Service{repository}
)

func Test_CreateCampaign(t *testing.T) {
	assertions := assert.New(t)

	t.Run("Save a new campaign", func(t *testing.T) {
		var campaign = contract.NewCampaignDto{
			Name:    "My campaign",
			Content: "Body of the campaign",
			Emails:  []string{"teste1@email.com"},
		}
		repository.On("Save", mock.MatchedBy(func(c *Campaign) bool {
			if c.Name != campaign.Name || c.Content != campaign.Content || len(c.Contacts) != len(campaign.Emails) {
				return false
			}
			return true
		})).Return(nil)

		createCampaign, err := service.CreateCampaign(campaign)
		if err != nil {
			return
		}

		assertions.NotNil(createCampaign)
		repository.AssertExpectations(t)
	})

	t.Run("Error trying to save a new campaign", func(t *testing.T) {
		var campaign = contract.NewCampaignDto{
			Name:    "",
			Content: "Body of the campaign",
			Emails:  []string{"teste1@email.com"},
		}
		repository.On("Save", mock.MatchedBy(func(c *Campaign) bool {
			if c.Name != campaign.Name || c.Content != campaign.Content || len(c.Contacts) != len(campaign.Emails) {
				return false
			}
			return true
		})).Return(nil)

		createCampaign, err := service.CreateCampaign(campaign)

		assertions.Empty(createCampaign)
		assertions.NotNil(err)

	})

	t.Run("Create a new campaign - empty name", func(t *testing.T) {
		campaign.Name = ""
		_, err := service.CreateCampaign(campaign)

		assertions.NotNil(err)
		assertions.Equal("name is less than the minimum 3", err.Error())
	})
}

func Test_CreateCampaign_ValidateRepository(t *testing.T) {
	assertions := assert.New(t)
	campaign = contract.NewCampaignDto{
		Name:    "My campaign",
		Content: "Body of the campaign",
		Emails: []string{
			"teste1@email.com",
		},
	}
	repository = new(repositoryMock)
	service = Service{repository}
	repository.On("Save", mock.Anything).Return(errors.New("error"))
	_, err := service.CreateCampaign(campaign)

	assertions.NotNil(err)
	assertions.Truef(errors.Is(err, internalerrors.ErrInternal), "expected internal server error, got %v", err)
	repository.AssertExpectations(t)
}
