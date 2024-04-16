package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) (int, error) {
	args := r.Called(campaign)

	// We'll need to do a bit of error checking to ensure the type assertion
	// doesn't cause a panic
	result, ok := args.Get(0).(int)
	if !ok {
		return 0, fmt.Errorf("error: return value is not int")
	}

	return result, args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	// args := r.Called(campaign)
	return nil, nil
}

// Cast para *Campaign
func (r *repositoryMock) GetBy(id int) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), nil
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
	service    = ServiceImpl{Repository: repository}
)

func Test_CreateCampaign(t *testing.T) {
	assertions := assert.New(t)

	t.Run("Save a new campaign", func(t *testing.T) {
		campaign := contract.NewCampaignDto{
			Name:    "My campaign",
			Content: "Body of the campaign",
			Emails:  []string{"teste1@email.com"},
		}
		repository.On("Save", mock.MatchedBy(func(c *Campaign) bool {
			if c.Name != campaign.Name || c.Content != campaign.Content ||
				len(c.Contacts) != len(campaign.Emails) {
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
		campaign := contract.NewCampaignDto{
			Name:    "",
			Content: "Body of the campaign",
			Emails:  []string{"teste1@email.com"},
		}
		repository.On("Save", mock.MatchedBy(func(c *Campaign) bool {
			if c.Name != campaign.Name || c.Content != campaign.Content ||
				len(c.Contacts) != len(campaign.Emails) {
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
	service = ServiceImpl{repository}
	repository.On("Save", mock.Anything).Return(errors.New("error"))
	_, err := service.CreateCampaign(campaign)

	assertions.NotNil(err)
	assertions.Truef(
		errors.Is(err, internalerrors.ErrInternal),
		"expected internal server error, got %v",
		err,
	)
	repository.AssertExpectations(t)
}

func Test_repositoryMock_Save(t *testing.T) {
	type args struct {
		campaign *Campaign
	}
	var tests []struct {
		name    string
		r       *repositoryMock
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.r.Save(tt.args.campaign); (err != nil) != tt.wantErr {
				t.Errorf("repositoryMock.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Get(t *testing.T) {
	var tests []struct {
		name    string
		r       *repositoryMock
		want    []Campaign
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("repositoryMock.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repositoryMock.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetBy(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(campaign.Name, campaign.Content, campaign.Emails)
	campaign.ID = 123
	repository := new(repositoryMock)
	repository.On("GetBy", mock.MatchedBy(func(id int) bool {
		return id == 123
	})).Return(campaign, nil)
	service.Repository = repository
	var campaignReturned, _ = service.GetBy(123)
	assertions.Equal(campaign.ID, campaignReturned.ID)
}

func Test_GetById_ReturnError(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(campaign.Name, campaign.Content, campaign.Emails)
	repository := new(repositoryMock)
	repository.On("GetBy", mock.Anything).Return(nil, errors.New("internal error"))
	service.Repository = repository
	var _, err = service.GetBy(campaign.ID)
	assertions.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
