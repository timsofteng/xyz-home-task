package googleBooks

import (
	"testing"
	"github.com/timsofteng/xyz-home-task/service"
)

func TestMapBookToServiceFormat(t *testing.T) {
	apiBook := Book{
		ID: "1",
		VolumeInfo: VolumeInfo{
			Title:       "Test Book",
			Description: "This is a test book.",
			PageCount:   300,
			ImageLinks:  ImageLinks{Thumbnail: "http://example.com/thumbnail.jpg"},
		},
		SaleInfo: SaleInfo{
			ListPrice: &Price{
				Amount:      9.99,
				CurrencyCode: "USD",
			},
		},
	}

	revision := 1

	expected := service.Book{
		ID:          apiBook.ID,
		Title:       apiBook.VolumeInfo.Title,
		Description: apiBook.VolumeInfo.Description,
		Pages:       apiBook.VolumeInfo.PageCount,
		Thumbnail:   apiBook.VolumeInfo.ImageLinks.Thumbnail,
		Price:       9.99,
		Currency:    "USD",
		Revision:    &revision,
	}

	result := mapBookToServiceFormat(apiBook, &revision)

	if result != expected {
		t.Errorf("expected %+v, got %+v", expected, result)
	}
}

func TestMapPrice(t *testing.T) {
	tests := []struct {
		price    *Price
		expected float32
	}{
		{&Price{Amount: 10.0}, 10.0},
		{nil, 0.0},
	}

	for _, test := range tests {
		result := mapPrice(test.price)
		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestMapCurrency(t *testing.T) {
	tests := []struct {
		price    *Price
		expected string
	}{
		{&Price{CurrencyCode: "USD"}, "USD"},
		{nil, "N/A"},
	}

	for _, test := range tests {
		result := mapCurrency(test.price)
		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

