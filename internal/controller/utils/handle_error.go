package utils

import (
	"github.com/go-chi/render"
	"net/http"
)

func HandleError500(err error) {
	render.Status(nil, http.StatusInternalServerError)
	render.JSON(nil, nil, map[string]string{"error": err.Error()})
}
