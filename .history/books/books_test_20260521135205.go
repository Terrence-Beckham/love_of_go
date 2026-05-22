package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	want := []books.Book{
		{
			Title:  "in the Company of Cheeful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		{
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
	got := books.GetAllBooks()
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}
func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
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
