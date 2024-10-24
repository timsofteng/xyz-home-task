package apperrors

import (
	"errors"
)

var (
	ErrOther             = errors.New("other")
	ErrAlreadyExist      = errors.New("already_exist")
	ErrNotFound          = errors.New("not_found")
	ErrInternal          = errors.New("internal_error")
	ErrExternal          = errors.New("external_error")
	ErrBadRequest        = errors.New("bad_request")
	ErrBadGateway        = errors.New("bad_gateway")
	ErrUnauthenticated   = errors.New("unauthenticated")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrPermissionDenied  = errors.New("permission_denied")
	ErrForbidden         = errors.New("forbidden")
	ErrValidation        = errors.New("validation")
	ErrTimeout           = errors.New("timeout")
	ErrUnavailable       = errors.New("unavailable")
	ErrNotImplemented    = errors.New("not_implemented")
	ErrResourceExhausted = errors.New("resource_exhausted")
	ErrCanceled          = errors.New("canceled")
	ErrConflict          = errors.New("conflict")
)
