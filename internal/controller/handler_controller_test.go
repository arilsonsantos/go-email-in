package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"emailn/internal/internalerrors"
	"github.com/stretchr/testify/assert"
)

func TestHandleControllerError(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   EndpointFunc
		wantStatus int
		wantBody   string
	}{
		{
			name: "Empty Response",
			endpoint: func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
				return nil, http.StatusNoContent, internalerrors.ErrNoContent
			},
			wantStatus: http.StatusNoContent,
			wantBody:   "{\"error\":\"no content\"}",
		},
		{
			name: "General Error",
			endpoint: func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
				return nil, 0, errors.New("error")
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "{\"error\":\"error\"}",
		},
		{
			name: "Internal Error",
			endpoint: func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
				return nil, 0, internalerrors.ErrInternal
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "{\"error\":\"internal server error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://test.com", nil)
			w := httptest.NewRecorder()

			HandleError(tt.endpoint)(w, req)

			resp := w.Result()
			body := w.Body.String()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
			assert.Equal(t, tt.wantBody, strings.TrimRight(body, "\n"))
		})
	}
}
