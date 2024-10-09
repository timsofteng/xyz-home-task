package e

import "errors"

func MapHTTPStatusCodeToInternalError(statusCode int) error {
	switch statusCode {
	case 400:
		return BadRequest
	case 401:
		return Unauthenticated
	case 403:
		return Unauthorized
	case 404:
		return NotFound
	case 409:
		return Conflict
	case 429:
		return ResourceExhausted
	case 500:
		return Internal
	case 502:
		return BadGateway
	case 504:
		return Timeout
	default:
		return External
	}
}

func MapInternalErrorToHTTPStatusCode(e error) int {
	switch {
	case errors.Is(e, BadRequest):
		return 400
	case errors.Is(e, Unauthenticated):
		return 401
	case errors.Is(e, Unauthorized):
		return 403
	case errors.Is(e, PermissionDenied):
		return 403
	case errors.Is(e, Forbidden):
		return 403
	case errors.Is(e, NotFound):
		return 404
	case errors.Is(e, Validation):
		return 422
	case errors.Is(e, Timeout):
		return 408
	case errors.Is(e, Unavailable):
		return 503
	case errors.Is(e, Internal):
		return 500
	case errors.Is(e, External):
		return 502
	case errors.Is(e, NotImplemented):
		return 501
	case errors.Is(e, ResourceExhausted):
		return 429
	case errors.Is(e, Canceled):
		return 499
	default:
		return 500 // Default to 500 for unknown internal errors
	}
}
