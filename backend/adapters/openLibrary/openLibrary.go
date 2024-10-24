package openlibrary

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/timsofteng/xyz-home-task/internal/errors"
	"github.com/timsofteng/xyz-home-task/internal/logger"
)

type BookInfo struct {
	Key      string `json:"key"`
	Revision int    `json:"revision,omitempty"`
}

type Client struct {
	logger  logger.Logger
	BaseURL string
	client  *http.Client
}

func New(logger logger.Logger) *Client {
	return &Client{
		logger:  logger.With("api", "open library"),
		BaseURL: "https://openlibrary.org",
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) getBookInfoByISBN(ctx context.Context, isbn string) (*BookInfo, error) {
	logger := c.logger.With("request", "get book by isbn")

	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet,
		fmt.Sprintf("%s/api/books?bibkeys=ISBN:%s&format=json&jscmd=data",
			c.BaseURL, isbn), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch book info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"failed to get book info: %w",
			apperrors.MapHTTPStatusCodeToInternalError(resp.StatusCode))
	}

	var result map[string]BookInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	logger.Debug("response", "body", result)

	bookInfo, ok := result[fmt.Sprintf("ISBN:%s", isbn)]
	if !ok {
		return nil, fmt.Errorf("book not found for ISBN: %s", isbn)
	}

	return &bookInfo, nil
}

func (c *Client) getRevision(ctx context.Context, key string) (*int, error) {
	logger := c.logger.With("request", "get revision by key")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s.json", c.BaseURL, key), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch revision: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Debug("request with bad code", "status", resp.StatusCode)
		return nil, fmt.Errorf(
			"failed to get revision: %w",
			apperrors.MapHTTPStatusCodeToInternalError(resp.StatusCode))
	}

	var bookInfo BookInfo
	if err := json.NewDecoder(resp.Body).Decode(&bookInfo); err != nil {
		return nil, err
	}

	return &bookInfo.Revision, nil
}

func (c *Client) GetRevisionByISBN(ctx context.Context, isbn string) (*int, error) {
	bookInfo, err := c.getBookInfoByISBN(ctx, isbn)
	if err != nil {
		return nil, err
	}

	revision, err := c.getRevision(ctx, bookInfo.Key)
	if err != nil {
		return nil, err
	}

	return revision, nil
}
