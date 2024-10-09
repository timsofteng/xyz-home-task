package e

import (
	"errors"
)

var (
	Other             = errors.New("other")
	AlreadyExist      = errors.New("already_exist")  
	NotFound          = errors.New("not_found")      
	Internal          = errors.New("internal_error") 
	External          = errors.New("external_error") 
	BadRequest        = errors.New("bad_request")
	BadGateway        = errors.New("bad_gateway")
	Unauthenticated   = errors.New("unauthenticated") 
	Unauthorized      = errors.New("unauthorized")    
	PermissionDenied  = errors.New("permission_denied")
	Forbidden         = errors.New("forbidden")
	Validation        = errors.New("validation")
	Timeout           = errors.New("timeout")
	Unavailable       = errors.New("unavailable")
	NotImplemented    = errors.New("not_implemented")
	ResourceExhausted = errors.New("resource_exhausted")
	Canceled          = errors.New("canceled")
	Conflict          = errors.New("conflict")
)
