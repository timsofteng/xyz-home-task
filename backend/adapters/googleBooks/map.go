package googleBooks

import "github.com/timsofteng/xyz-home-task/service"

func mapBookToServiceFormat(apiBook Book, rev *int) service.Book {
	return service.Book{
		ID:          apiBook.ID,
		Title:       apiBook.VolumeInfo.Title,
		Description: apiBook.VolumeInfo.Description,
		Pages:       apiBook.VolumeInfo.PageCount,
		Thumbnail:   apiBook.VolumeInfo.ImageLinks.Thumbnail,
		Price:       mapPrice(apiBook.SaleInfo.ListPrice),
		Currency:    mapCurrency(apiBook.SaleInfo.ListPrice),
		Revision:    rev,
	}
}

// mapPrice extracts the price from the SaleInfo.
func mapPrice(price *Price) float32 {
	if price != nil {
		return price.Amount
	}
	return 0.0
}

// mapCurrency extracts the currency from the SaleInfo.
func mapCurrency(price *Price) string {
	if price != nil {
		return price.CurrencyCode
	}
	return "N/A"
}
