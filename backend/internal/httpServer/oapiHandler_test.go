package httpServer

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timsofteng/xyz-home-task/internal/logger"
)

var mockLogger = logger.New("debug")

func TestValidationErrorHandler(t *testing.T) {
	handler := validationErrorHandler(mockLogger)

	rr := httptest.NewRecorder()

	handler(rr, "validation error", http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var resp Error
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "validation error", resp.Message)
}
