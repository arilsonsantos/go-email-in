package controller

import (
	"emailn/internal/internalerrors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_endpoints_return_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}

	handerfunc := HandleError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handerfunc.ServeHTTP(res, req)
	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_HandlerError_when_endpoints_return_business_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("business error")
	}

	handerfunc := HandleError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handerfunc.ServeHTTP(res, req)
	assert.Equal(http.StatusBadRequest, res.Code)
}

func Test_HandlerError_when_endpoints_return_object_success(t *testing.T) {
	{
		assert := assert.New(t)

		type bodyExpeted struct {
			Id int
		}
		objExpected := bodyExpeted{Id: 10}
		endpoint := func(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
			return objExpected, 201, nil

		}

		handerfunc := HandleError(endpoint)
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		handerfunc.ServeHTTP(res, req)

		assert.Equal(http.StatusCreated, res.Code)
		objReturned := bodyExpeted{}
		json.Unmarshal(res.Body.Bytes(), &objReturned)
		assert.Equal(objExpected, objReturned)
	}
}
