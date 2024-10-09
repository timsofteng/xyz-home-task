package googleBooks

import (
	"sort"
	"testing"

	"github.com/timsofteng/xyz-home-task/service"
)

func TestSortByTitle(t *testing.T) {
	books := []service.Book{
		{Title: "B Book"},
		{Title: "A Book"},
		{Title: "C Book"},
	}

	expected := []service.Book{
		{Title: "A Book"},
		{Title: "B Book"},
		{Title: "C Book"},
	}

	sort.Sort(byTitle(books))

	for i := range expected {
		if books[i].Title != expected[i].Title {
			t.Errorf("expected %v, got %v", expected[i].Title, books[i].Title)
		}
	}
}
