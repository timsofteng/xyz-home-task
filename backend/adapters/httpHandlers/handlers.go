package httpHandlers

import (
	"context"
	"time"

	"github.com/timsofteng/xyz-home-task/internal/logger"
	"github.com/timsofteng/xyz-home-task/internal/httpServer"
	"github.com/timsofteng/xyz-home-task/service"
)

type handlers struct {
	logger       logger.Logger
	booksService *service.BooksService
}

func New(
	logger logger.Logger,
	booksService *service.BooksService) *handlers {
	return &handlers{
		logger: logger, booksService: booksService,
	}
}

func (h *handlers) GetBooks(
	ctx context.Context,
	request httpServer.GetBooksRequestObject,
) (httpServer.GetBooksResponseObject, error) {
	logger := h.logger.With("http handler", "get books")
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	logger.Info("request from client", "requestData", request)

	books, err := h.booksService.GetBooks(ctx, request.Params.Q)

	// Check if the context was canceled due to a timeout
	if ctx.Err() == context.DeadlineExceeded {
		logger.Error("request timed out")
		return httpServer.GetBooks408JSONResponse{Message: "Timeout"}, nil
	}

	if err != nil {
		logger.Error("failed to get books from service", "error", err)
		return httpServer.GetBooks500JSONResponse{Message: "Internal error"}, nil
	}

	items := mapServiceBooksToHTTPBooks(books)

	logger.Info("sending response", "items", items)

	return httpServer.GetBooks200JSONResponse{
		Items: items,
	}, nil
}
