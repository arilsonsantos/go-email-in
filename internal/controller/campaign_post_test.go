package controller

import (
	"bytes"
	"context"
	"emailn/internal/controller/dto"
	"emailn/internal/internalerrors"
	"emailn/internal/test/mocks"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignPost_should_save_campaing(t *testing.T) {
	assertions := assert.New(t)
	service := new(mocks.CampaignServiceMock)
	createdByEmail := "teste@email.com"
	body := dto.NewPostCampaignDto{
		Name:    "My campaign",
		Content: "Body of the campaign",
		Emails:  []string{"teste@example.com"},
	}
	service.On("CreateCampaign", mock.MatchedBy(func(dto dto.NewPostCampaignDto) bool {
		if dto.Name == body.Name && dto.Content == body.Content &&
			len(dto.Emails) == len(body.Emails) &&
			dto.CreatedBy == createdByEmail {
			return true
		}
		return false
	})).Return(123, nil)

	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return
	}
	req, _ := http.NewRequest("POST", "/campaign", &buf)
	ctx := context.WithValue(req.Context(), "email", createdByEmail)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	obj, status, err := handler.CampaignPost(rr, req)
	assertions.Equal(http.StatusCreated, status)
	assertions.Nil(err)
	assertions.Equal(map[string]int{"id": 123}, obj)
	assertions.Equal(123, obj.(map[string]int)["id"])
}

func Test_CampaignPost_should_return_error(t *testing.T) {
	assertions := assert.New(t)
	service := new(mocks.CampaignServiceMock)
	createdByEmail := "teste@email.com"
	body := dto.NewPostCampaignDto{
		Name:    "My campaign",
		Content: "Body of the campaign",
		Emails:  []string{"teste@example.com"},
	}
	service.On("CreateCampaign", mock.Anything).Return(0, fmt.Errorf("internal server error"))

	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return
	}
	req, _ := http.NewRequest("POST", "/campaign", &buf)
	ctx := context.WithValue(req.Context(), "email", createdByEmail)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)
	assertions.NotNil(err)
	assertions.Equal(http.StatusCreated, status)
	assertions.Equal(err.Error(), internalerrors.ErrInternal.Error())
}
