package books_test

import (
	"books"
	"slices"
	"testing"

)

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	want := []books.Book{
		{ID: 1,
			Title:  "The Mandalorian",
			Author: "George Lucas",
			Copies: 4,
		},
		{ID: 2,
			Title:  "Darth Bane",
			Author: "Alex Karpashy ",
			Copies: 2,
		},
		{ID: 3,
			Title:  "The Rule of Two",
			Author: "Hurst",
			Copies: 9,
		},
		{ID: 4,
			Title:  "Legacy of Evil",
			Author: "Someone",
			Copies: 3,
		},
	}
	got := books.GetAllBooks()
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}
func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		ID:     1,
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := books.BookToString(input)

	if want != got {
		t.Fatalf("%q not equal to %q", want, got)
	}

}

func TestGetBook_ReturnCorrectBook(t *testing.T) {
	t.Parallel()
		want := books.Book{
		ID:     4,
		Title:  "Legacy of Evil",
		Author: "Someone",
		Copies: 3,
	}
	got ,ok := books.GetBook(4)
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}
