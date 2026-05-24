package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func GetTestCatalog() books.Catalog {
	return books.Catalog{
		"def": {
			ID:     "def",
			Title:  "Darth Bane",
			Author: "Alex Karpashy ",
			Copies: 2,
		},
		"abc": {
			ID:     "abc",
			Title:  "The Mandalorian",
			Author: "George Lucas",
			Copies: 4,
		},
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	var catalog = GetTestCatalog()
	want := []books.Book{
		{
			ID:     "def",
			Title:  "Darth Bane",
			Author: "Alex Karpashy ",
			Copies: 2,
		},
		{
			ID:     "abc",
			Title:  "The Mandalorian",
			Author: "George Lucas",
			Copies: 4,
		},
	}

	got := catalog.GetAllBooks()
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}
func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		ID:     "abc",
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := input.BookToString()

	if want != got {
		t.Fatalf("%q not equal to %q", want, got)
	}

}

func TestGetBook_ReturnCorrectBook(t *testing.T) {
	t.Parallel()

	var catalog = GetTestCatalog()
	want := books.Book{

		ID:     "abc",
		Title:  "The Mandalorian",
		Author: "George Lucas",
		Copies: 4,
	}
	got, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("Book not found")
	}

	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}
func TestGetBook_ReturnsFalseIfNotFound(t *testing.T) {
	t.Parallel()
	var catalog = GetTestCatalog()
	_, ok := catalog.GetBook("sdfs")
	if ok {
		t.Fatal("Nonexistent ID")
	}

	if ok {
		t.Fatal("Want false for nonexistent ID got true")
	}

}

func TestAddBook(t *testing.T) {
	t.Parallel()

	var catalog = GetTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("The book already exists")
	}
	catalog.AddBook(books.Book{
		ID:     "123",
		Title:  "New Glory",
		Author: "Me Again",
		Copies: 25,
	})

	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}
func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T){
	t.Parallel()
	book := books.Book{
		Copies:5,
	}
	book.SetCopies(12)
	if book.Copies != 12{
		t.Errorf()
	}
}