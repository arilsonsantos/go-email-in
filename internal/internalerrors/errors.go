package internalerrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrNoContent = errors.New("no content")
