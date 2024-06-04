package controller

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Auth_When_Authorization_Is_Missing_Return_Error(t *testing.T) {
	assertion := assert.New(t)
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")

	})
	handlerFunc := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)
	assertion.Equal(http.StatusUnauthorized, res.Code)
	response, _ := io.ReadAll(res.Body)
	var responseMap map[string]string
	_ = json.Unmarshal(response, &responseMap)
	assertion.Equal("request does not contain an authorization token", responseMap["error"])
}

func Test_Auth_When_Authorization_Is_Invalid_Return_Error(t *testing.T) {
	assertion := assert.New(t)
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")

	})
	ValidationToken = func(tokenStr string) (string, error) {
		return "", errors.New("invalid token")
	}
	handlerFunc := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer xpto")
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)
	assertion.NotNil(ValidationToken)
	assertion.Equal(http.StatusUnauthorized, res.Code)
	assertion.Contains(res.Body.String(), "invalid token")
}

func Test_Auth_When_Authorization_Is_Valid_Call_Next_Handler(t *testing.T) {
	assertion := assert.New(t)
	var email string
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email = r.Context().Value("email").(string)
	})

	ValidationToken = func(tokenStr string) (string, error) {
		return "test@email.com", nil
	}
	handlerFunc := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer xptoToken")
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)
	assertion.NotNil(ValidationToken)
	assertion.Equal(http.StatusOK, res.Code)
	assertion.Equal(email, "test@email.com")
}
