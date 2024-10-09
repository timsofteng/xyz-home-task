package httpHandlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timsofteng/xyz-home-task/internal/httpServer"
	"github.com/timsofteng/xyz-home-task/service"
)

func TestMapServiceBookToHTTPBook(t *testing.T) {
	serviceBook := service.Book{
		ID:          "123",
		Title:       "Test Book",
		Price:       19.99,
		Description: "A great book",
		Pages:       300,
		Currency:    "USD",
		Thumbnail:   "http://example.com/thumbnail.jpg",
		Revision:    nil,
	}

	expectedHTTPBook := httpServer.BookBase{
		Id:          "123",
		Title:       "Test Book",
		Price:       19.99,
		Description: "A great book",
		Pages:       300,
		Currency:    "USD",
		Thumbnail:   "http://example.com/thumbnail.jpg",
		Revision:    nil,
	}

	httpBook := mapServiceBookToHTTPBook(serviceBook)
	assert.Equal(t, expectedHTTPBook, httpBook)
}

func TestMapServiceBooksToHTTPBooks(t *testing.T) {
	serviceBooks := []service.Book{
		{
			ID:          "123",
			Title:       "Test Book 1",
			Price:       19.99,
			Description: "A great book",
			Pages:       300,
			Currency:    "USD",
			Thumbnail:   "http://example.com/thumbnail1.jpg",
			Revision:    nil,
		},
		{
			ID:          "456",
			Title:       "Test Book 2",
			Price:       29.99,
			Description: "Another great book",
			Pages:       400,
			Currency:    "EUR",
			Thumbnail:   "http://example.com/thumbnail2.jpg",
			Revision:    intPointer(3),
		},
	}

	expectedHTTPBooks := []httpServer.BookBase{
		{
			Id:          "123",
			Title:       "Test Book 1",
			Price:       19.99,
			Description: "A great book",
			Pages:       300,
			Currency:    "USD",
			Thumbnail:   "http://example.com/thumbnail1.jpg",
			Revision:    nil,
		},
		{
			Id:          "456",
			Title:       "Test Book 2",
			Price:       29.99,
			Description: "Another great book",
			Pages:       400,
			Currency:    "EUR",
			Thumbnail:   "http://example.com/thumbnail2.jpg",
			Revision:    intPointer(3),
		},
	}

	httpBooks := mapServiceBooksToHTTPBooks(serviceBooks)
	assert.Equal(t, expectedHTTPBooks, httpBooks)
}

func TestMapRevision(t *testing.T) {
	t.Run("Nil revision", func(t *testing.T) {
		var rev *int = nil
		assert.Nil(t, mapRevision(rev))
	})

	t.Run("Revision less than 2", func(t *testing.T) {
		rev := 1
		assert.Nil(t, mapRevision(&rev))
	})

	t.Run("Revision greater than or equal to 2", func(t *testing.T) {
		rev := 3
		result := mapRevision(&rev)
		assert.NotNil(t, result)
		assert.Equal(t, rev, *result)
	})
}

// Helper function for making pointers to int
func intPointer(i int) *int {
	return &i
}
