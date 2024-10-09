package httpServer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/timsofteng/xyz-home-task/internal/e"
	"github.com/timsofteng/xyz-home-task/internal/logger"
)

type Error struct {
	Message string `json:"message"`
}

func validationErrorHandler(logger logger.Logger) middleware.ErrorHandler {
	return func(
		w http.ResponseWriter, message string, statusCode int,
	) {
		logger := logger.With("handler", "validation error handler")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		resp := Error{
			Message: message,
		}

		logger.Error(message)

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			logger.Error("convert to json",
				"details", err.Error(),
			)
		}
	}
}

func responseErrorHandler(
	w http.ResponseWriter, r *http.Request, err error,
) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(e.MapInternalErrorToHTTPStatusCode(err))

	encodeErr := json.NewEncoder(w).Encode(err)
	if encodeErr != nil {
		log.Printf("Error encoding response: %v", encodeErr)
	}
}

func WrapToOapiHandler(
	logger logger.Logger,
	mux *http.ServeMux,
	handlers StrictServerInterface,
) (http.Handler, error) {

	swagger, err := GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}

	strictHandlers := NewStrictHandlerWithOptions(
		handlers, nil,
		StrictHTTPServerOptions{
			ResponseErrorHandlerFunc: responseErrorHandler,
		},
	)

	// maybe it make sense to create separate middleware
	// for base url
	HandlerFromMuxWithBaseURL(strictHandlers, mux, "/api/v1")

	h := middleware.OapiRequestValidatorWithOptions(
		swagger,
		&middleware.Options{ErrorHandler: validationErrorHandler(logger)},
	)(mux)

	return h, nil
}
