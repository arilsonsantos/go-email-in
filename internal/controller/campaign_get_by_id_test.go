package controller

import (
	"emailn/internal/controller/dto"
	"emailn/internal/domain/campaign"
	"emailn/internal/test/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CampaignGetById_should_return_campaing(t *testing.T) {
	assertions := assert.New(t)
	campaignResponse := dto.NewGetCampaignDto{
		ID:      123,
		Name:    "My campaign",
		Content: "Body of the campaign",
		Status:  campaign.Pending,
	}
	service := new(mocks.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaignResponse, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	obj, status, err := handler.CampaignGetById(rr, req)
	assertions.Equal(http.StatusOK, status)
	assertions.Nil(err)
	assertions.Equal(campaignResponse.ID, obj.(*dto.NewGetCampaignDto).ID)
}

func TestHandler_CampaignGetByIdPost_should_return_error(t *testing.T) {
	assertions := assert.New(t)
	service := new(mocks.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(nil, errors.New("internal server error"))
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignGetById(rr, req)
	assertions.Equal(http.StatusNoContent, status)
	assertions.NotNil(err)
}
