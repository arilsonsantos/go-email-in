package utils

import (
	"emailn/internal/internalerrors"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

func HandleError(w http.ResponseWriter, req *http.Request, err error) {
	if errors.Is(err, internalerrors.ErrInternal) {
		render.Status(req, http.StatusInternalServerError)
	} else {
		render.Status(req, http.StatusBadRequest)
	}
	render.JSON(w, req, map[string]string{"error": err.Error()})
}

func HandleError500(err error) {
	render.Status(nil, http.StatusInternalServerError)
	render.JSON(nil, nil, map[string]string{"error": err.Error()})
}
