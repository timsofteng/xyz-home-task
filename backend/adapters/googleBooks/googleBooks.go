package googleBooks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/timsofteng/xyz-home-task/internal/errors"
	"github.com/timsofteng/xyz-home-task/internal/logger"
	"github.com/timsofteng/xyz-home-task/service"
)

// Repo represents the Google Books API repository.
type Repo struct {
	baseURL      string
	httpClient   *http.Client
	logger       logger.Logger
	revisionRepo RevisionRepo
}

// GoogleBooksAPIResponse represents the response structure from the Google Books API.
type GoogleBooksAPIResponse struct {
	Kind       string `json:"kind" validate:"required"`
	TotalItems int    `json:"totalItems"`
	Items      []Book `json:"items"`
}

// Book represents the structure of each book in the Google Books API response.
type Book struct {
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	ID         string     `json:"id"`
}

// VolumeInfo holds the book's metadata from the Google Books API.
type VolumeInfo struct {
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	PageCount           int                  `json:"pageCount"`
	ImageLinks          ImageLinks           `json:"imageLinks"`
	MaturityRating      string               `json:"maturityRating"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
}

// SaleInfo holds the book's sale details, including price and currency.
type SaleInfo struct {
	ListPrice  *Price `json:"listPrice"`
	Salebility string `json:"salebility"`
}

// Price holds the price and currency information of a book.
type Price struct {
	Amount       float32 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}

// Price holds the price and currency information of a book.
type ImageLinks struct {
	Thumbnail string `json:"smallThumbnail"`
}

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type RevisionRepo interface {
	GetRevisionByISBN(ctx context.Context, isbn string) (*int, error)
}

// New creates a new Google Books repository.
func New(logger logger.Logger, ri RevisionRepo) *Repo {
	return &Repo{
		baseURL:      "https://www.googleapis.com/books/v1/volumes",
		httpClient:   &http.Client{Timeout: time.Second * 5},
		logger:       logger.With("api", "google book library"),
		revisionRepo: ri,
	}
}

// GetBooks fetches books based on a search query from the Google Books API.
func (r *Repo) GetBooks(ctx context.Context, query string) ([]service.Book, error) {
	logger := r.logger.With("request", "get books")

	url := fmt.Sprintf("%s?q=%s", r.baseURL, query)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}
	defer res.Body.Close()

	var apiResponse GoogleBooksAPIResponse
	if err := json.NewDecoder(res.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if res.StatusCode != 200 {
		logger.Debug("bad code", "status", res.StatusCode)
		return nil, fmt.Errorf(
			"error from api: %w",
			apperrors.MapHTTPStatusCodeToInternalError(res.StatusCode))
	}

	logger.Debug("resp", "body", apiResponse)

	books, err := r.toServiceFormat(ctx, apiResponse.Items)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repo) toServiceFormat(ctx context.Context, apiBooks []Book) ([]service.Book, error) {
	var (
		books  []service.Book
		wg     sync.WaitGroup
		mu     sync.Mutex
		bookCh = make(chan service.Book)
	)

	// Collect results from the channel in a separate goroutine
	go func() {
		for book := range bookCh {
			mu.Lock()
			books = append(books, book)
			mu.Unlock()
		}
	}()

	for _, apiBook := range apiBooks {
		apiBook := apiBook

		if apiBook.VolumeInfo.MaturityRating != "NOT_MATURE" || apiBook.SaleInfo.Salebility == "NOT_FOR_SALE" {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			isbn := findIsbnInIdettifiers(apiBook.VolumeInfo.IndustryIdentifiers)

			select {
			case <-ctx.Done():
				r.logger.Debug("timeout occurred", "isbn", isbn)
				return
			default:
				var rev *int
				var err error
				if len(isbn) > 0 {
					rev, err = r.revisionRepo.GetRevisionByISBN(ctx, isbn)
					if err != nil {
						r.logger.Debug("failed to get revision", "isbn", isbn, "details", err)
						return
					}
				}
				book := mapBookToServiceFormat(apiBook, rev)
				bookCh <- book
			}
		}()
	}

	wg.Wait()

	close(bookCh)

	sort.Sort(byTitle(books))

	return books, nil
}

func findIsbnInIdettifiers(ii []IndustryIdentifier) string {
	if len(ii) < 1 {
		return ""
	}

	re := regexp.MustCompile("isbn")

	for _, i := range ii {
		t := strings.ToLower(i.Type)

		if re.MatchString(t) {
			return i.Identifier
		}
	}

	return ""
}
