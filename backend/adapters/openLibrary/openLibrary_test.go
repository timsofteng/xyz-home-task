package openlibrary

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timsofteng/xyz-home-task/internal/logger"
)

var MockLogger = logger.New("debug")

func TestGetBookInfoByISBN(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/books" && r.Method == http.MethodGet {
			response := `{"ISBN:1234567890": {"key": "/books/1234567890", "revision": 1}}`
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(response))
			if err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	client := New(MockLogger)
	client.BaseURL = ts.URL

	bookInfo, err := client.getBookInfoByISBN(context.Background(), "1234567890")

	assert.NoError(t, err)
	assert.NotNil(t, bookInfo)
	assert.Equal(t, "/books/1234567890", bookInfo.Key)
	assert.Equal(t, 1, bookInfo.Revision)
}

func TestGetRevision(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/books/1234567890.json" && r.Method == http.MethodGet {
			response := `{"key": "/books/1234567890", "revision": 2}`
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(response))
			if err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	client := New(MockLogger)
	client.BaseURL = ts.URL

	revision, err := client.getRevision(context.Background(), "/books/1234567890")

	assert.NoError(t, err)
	assert.NotNil(t, revision)
	assert.Equal(t, 2, *revision)
}

func TestGetRevisionByISBN(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/books" && r.Method == http.MethodGet {
			response := `{"ISBN:1234567890": {"key": "/books/1234567890", "revision": 1}}`
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(response))
			if err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			return
		} else if r.URL.Path == "/books/1234567890.json" && r.Method == http.MethodGet {
			response := `{"key": "/books/1234567890", "revision": 2}`
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(response))
			if err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	client := New(MockLogger)
	client.BaseURL = ts.URL

	revision, err := client.GetRevisionByISBN(context.Background(), "1234567890")

	assert.NoError(t, err)
	assert.NotNil(t, revision)
	assert.Equal(t, 2, *revision)
}

func TestGetBookInfoError(t *testing.T) {
	client := New(MockLogger)
	client.BaseURL = "http://invalid.url"

	_, err := client.getBookInfoByISBN(context.Background(), "1234567890")

	assert.Error(t, err)
	// need to check for a specific custom error
	assert.Contains(t, err.Error(), "failed to fetch book info")
}
