package campaign

import (
	"context"
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(ctx context.Context, campaign *Campaign) (int, error) {
	args := r.Called(campaign)

	result, ok := args.Get(0).(int)
	if !ok {
		return 0, fmt.Errorf("error: return value is not int")
	}

	return result, args.Error(1)
}

func (r *repositoryMock) Get() (*[]Campaign, error) {
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

func (r *repositoryMock) Update(campaign *Campaign) error {
	return nil
}

var (
	campaign = contract.NewPostCampaignDto{
		Name:    "My campaign",
		Content: "Body of the campaign",
		Emails: []string{
			"teste1@email.com",
		},
		CreatedBy: "teste@email.com",
	}
	campaignPending *Campaign
	repository      *repositoryMock
	service         = ServiceImpl{}
)

func setUp() {
	campaignPending, _ = NewCampaign(campaign.Name, campaign.Content, campaign.Emails, campaign.CreatedBy)
	repository = new(repositoryMock)
	service.Repository = repository
}

func Test_CreateCampaign(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	ctx := context.Background()

	t.Run("Save a new campaign", func(t *testing.T) {
		campaign := contract.NewPostCampaignDto{
			Name:      "My campaign",
			Content:   "Body of the campaign",
			Emails:    []string{"teste1@email.com"},
			CreatedBy: "teste@email.com",
		}
		repository.On("Save", mock.Anything).Return(1, nil)

		ctx := context.Background()
		campaignId, err := service.CreateCampaign(ctx, campaign)
		if err != nil {
			return
		}

		assertions.NotNil(campaignId)
		repository.AssertExpectations(t)
	})

	t.Run("Error trying to save a new campaign", func(t *testing.T) {
		campaign := contract.NewPostCampaignDto{
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

		createCampaign, err := service.CreateCampaign(ctx, campaign)

		assertions.Empty(createCampaign)
		assertions.NotNil(err)
	})

	t.Run("Create a new campaign - empty name", func(t *testing.T) {
		campaign.Name = ""
		_, err := service.CreateCampaign(ctx, campaign)

		assertions.NotNil(err)
		assertions.Equal("name is less than the minimum 3", err.Error())
	})
}

func Test_CreateCampaign_ValidateRepository(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	ctx := context.Background()

	campaign = contract.NewPostCampaignDto{
		Name:    "My campaign",
		Content: "Body of the campaign",
		Emails: []string{
			"teste1@email.com",
		},
		CreatedBy: "teste@email.com",
	}
	repository.On("Save", mock.Anything).Return(errors.New("error"))
	_, err := service.CreateCampaign(ctx, campaign)

	assertions.NotNil(err)
	assertions.Truef(
		errors.Is(err, internalerrors.ErrInternal),
		"expected internal server error, got %v",
		err,
	)
	repository.AssertExpectations(t)
}

func Test_repositoryMock_Save(t *testing.T) {
	ctx := context.Background()
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
			if _, err := tt.r.Save(ctx, tt.args.campaign); (err != nil) != tt.wantErr {
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
	setUp()
	assertions := assert.New(t)
	campaignPending.ID = 123

	repository.On("GetBy", mock.MatchedBy(func(id int) bool {
		return id == 123
	})).Return(campaignPending, nil)

	var campaignReturned, _ = service.GetBy(123)
	assertions.Equal(campaignPending.ID, campaignReturned.ID)
	assertions.Equal("teste@email.com", campaign.CreatedBy)
}

func Test_GetById_ReturnError(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	repository.On("GetBy", mock.Anything).Return(nil, errors.New("internal error"))
	var _, err = service.GetBy(0)
	assertions.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Start_ReturnNoFound_When_Campaing_Does_Not_Exist(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	repository.On("GetBy", mock.Anything).Return(nil, errors.New("not found"))
	err := service.Start(0)
	assertions.Equal(internalerrors.ErrNotFound.Error(), err.Error())
}

func Test_Start_ReturnError_When_Campaing_Has_Status_Not_Equal_Pendent(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	campaignPending.Started()
	repository.On("GetBy", mock.Anything).Return(campaignPending, nil)
	err := service.Start(campaign.ID)
	assertions.Equal(internalerrors.ErrInvalidStatus.Error(), err.Error())
}

func Test_Start_Should_Send_Email(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	repository.On("GetBy", mock.Anything).Return(campaignPending, nil)

	emailWasSent := false
	sendEmail := func(campaign *Campaign) error {
		if campaign.ID == campaignPending.ID {
			emailWasSent = true
		}
		return nil
	}

	service.SendEmail = sendEmail
	_ = service.Start(campaignPending.ID)
	time.Sleep(1 * time.Second)
	assertions.True(emailWasSent)
}

func Test_Start_Return_Error_When_SendMail_Fail(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	repository.On("GetBy", mock.Anything).Return(campaignPending, nil)

	sendEmail := func(campaign *Campaign) error {
		return errors.New("error on send email")
	}

	service.SendEmail = sendEmail

	service.Start(campaignPending.ID)
	time.Sleep(2 * time.Second)
	assertions.Equal(Failed, campaignPending.Status)
}

func Test_Start_Return_Error_When_SendMail_Started(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	repository.On("GetBy", mock.Anything).Return(campaignPending, nil)

	sendEmail := func(campaign *Campaign) error {
		return errors.New("error on send email")
	}

	service.SendEmail = sendEmail

	_ = service.Start(campaignPending.ID)

	assertions.Equal(Started, campaignPending.Status)
}

func Test_Start_Return_Nil_When_Update_To_Done(t *testing.T) {
	setUp()
	assertions := assert.New(t)
	repository.On("GetBy", mock.Anything).Return(campaignPending, nil)
	repository.On("Update", mock.MatchedBy(func(id int) bool {
		return campaignPending.ID == id
	})).Return(nil)

	sendEmail := func(campaign *Campaign) error {
		return nil
	}

	service.SendEmail = sendEmail
	err := service.Start(campaignPending.ID)
	if err != nil {
		return
	}
	assertions.Equal(Started, campaignPending.Status)
	time.Sleep(1 * time.Second)
	assertions.Equal(Done, campaignPending.Status)
}
