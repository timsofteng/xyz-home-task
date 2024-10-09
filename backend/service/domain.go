package service

type Book struct {
	// Unique ID for the book
	ID string

	// Title of the book
	Title string

	// Description of the book
	Description string

	// Pages in book
	Pages int

	// Price of the book
	Price float32

	// Currency
	Currency string

	// Thumbnail URL
	Thumbnail string

	Revision *int
}
