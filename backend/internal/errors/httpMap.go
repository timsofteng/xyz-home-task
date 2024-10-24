package apperrors

import "errors"

func MapHTTPStatusCodeToInternalError(statusCode int) error {
	switch statusCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrUnauthenticated
	case 403:
		return ErrUnauthorized
	case 404:
		return ErrNotFound
	case 409:
		return ErrConflict
	case 429:
		return ErrResourceExhausted
	case 500:
		return ErrInternal
	case 502:
		return ErrBadGateway
	case 504:
		return ErrTimeout
	default:
		return ErrExternal
	}
}

func MapInternalErrorToHTTPStatusCode(e error) int {
	switch {
	case errors.Is(e, ErrBadRequest):
		return 400
	case errors.Is(e, ErrUnauthenticated):
		return 401
	case errors.Is(e, ErrUnauthorized):
		return 403
	case errors.Is(e, ErrPermissionDenied):
		return 403
	case errors.Is(e, ErrForbidden):
		return 403
	case errors.Is(e, ErrNotFound):
		return 404
	case errors.Is(e, ErrValidation):
		return 422
	case errors.Is(e, ErrTimeout):
		return 408
	case errors.Is(e, ErrUnavailable):
		return 503
	case errors.Is(e, ErrInternal):
		return 500
	case errors.Is(e, ErrExternal):
		return 502
	case errors.Is(e, ErrNotImplemented):
		return 501
	case errors.Is(e, ErrResourceExhausted):
		return 429
	case errors.Is(e, ErrCanceled):
		return 499
	default:
		return 500 // Default to 500 for unknown internal errors
	}
}
