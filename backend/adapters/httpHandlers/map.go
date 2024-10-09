package httpHandlers

import (
	"github.com/timsofteng/xyz-home-task/internal/httpServer"
	"github.com/timsofteng/xyz-home-task/service"
)

func mapServiceBookToHTTPBook(b service.Book) httpServer.BookBase {

	return httpServer.BookBase{
		Id:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Pages:       b.Pages,
		Currency:    b.Currency,
		Thumbnail:   b.Thumbnail,
		Revision:    mapRevision(b.Revision),
	}
}

func mapServiceBooksToHTTPBooks(b []service.Book) []httpServer.BookBase {
	httpBooks := make([]httpServer.BookBase, len(b))
	for i, svcAccount := range b {
		httpBooks[i] = mapServiceBookToHTTPBook(svcAccount)
	}
	return httpBooks
}

func mapRevision(r *int) *int {
	if r == nil {
		return r
	}

	if *r < 2 {
		return nil
	}

	return r

}
