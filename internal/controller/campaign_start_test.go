package controller

import (
	"context"
	"emailn/internal/test/mocks"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler_CampaignStart_should_send_email(t *testing.T) {
	assertions := assert.New(t)
	service := new(mocks.CampaignServiceMock)
	campaignId := 123
	service.On("Start", mock.MatchedBy(func(id int) bool {
		return id == campaignId
	})).Return(nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", strconv.Itoa(campaignId))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)
	assertions.Equal(http.StatusOK, status)
	assertions.Nil(err)
}

func TestHandler_CampaignStart_should_return_error(t *testing.T) {
	assertions := assert.New(t)
	service := new(mocks.CampaignServiceMock)
	service.On("Start", mock.Anything).Return(errors.New("internal server error"))
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)
	assertions.Equal(http.StatusUnprocessableEntity, status)
	assertions.NotNil(err)
}
