package utils

import (
	"github.com/go-chi/render"
	"net/http"
)

func ParseJSONRequest[T interface{}](req *http.Request, request T) (T, error) {
	err := render.DecodeJSON(req.Body, &request)
	return request, err
}
