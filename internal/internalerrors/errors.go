package internalerrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrNoContent = errors.New("no content")
var ErrNotFound = errors.New("not found")
var ErrInvalidStatus = errors.New("invalid status")
var ErrSendingEmail = errors.New("error sending email")
