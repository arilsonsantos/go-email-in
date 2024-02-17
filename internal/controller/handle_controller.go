package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandleControllerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		obj, status, err := endpointFunc(w, r)
		if err != nil {
			render.Status(r, status)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	}
}
