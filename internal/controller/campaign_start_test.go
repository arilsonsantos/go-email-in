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
	service := new(mocks.CampaignServiceMock)
	campaignId := 123
	service.On("Start", mock.MatchedBy(func(id int) bool {
		return id == campaignId
	})).Return(nil)
	handler := Handler{CampaignService: service}
	req, _, chiContext := newHttpTest(campaignId, "PATCH")
	req = addParameter(req, chiContext)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)
	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)
}

func TestHandler_CampaignStart_should_return_error(t *testing.T) {
	service := new(mocks.CampaignServiceMock)
	service.On("Start", mock.Anything).Return(errors.New("internal server error"))
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, status)
	assert.NotNil(t, err)
}

func addParameter(req *http.Request, chiContext *chi.Context) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
}

func newHttpTest(campaignId int, method string) (*http.Request, error, *chi.Context) {
	req, _ := http.NewRequest(method, "/", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", strconv.Itoa(campaignId))
	return req, nil, chiContext
}
