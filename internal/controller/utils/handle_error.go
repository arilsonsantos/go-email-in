package utils

import (
	"emailn/internal/internalerrors"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

func HandleError(err error) (int, error) {
	if errors.Is(err, internalerrors.ErrInternal) {
		return http.StatusInternalServerError, err
	} else {
		return http.StatusBadRequest, err
	}
}

func HandleError500(err error) {
	render.Status(nil, http.StatusInternalServerError)
	render.JSON(nil, nil, map[string]string{"error": err.Error()})
}
