package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	want := []books.Book{
	{
		Title: "The Mandalorian",
		Author: "George Lucas",
		Copies: 4,
	},
		{
		Title: "Darth Bane",
		Author: "Alex Karpashy ",
		Copies: 2,
	},
		{
		Title: "The Rule of Two",
		Author: "Hurst",
		Copies: 9,
	},
		{
		Title: "Legacy of Evil",
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

func TestGetBook_ReturnCorrectBook(t *testing.T)  {
	t.Parallel()
	
	
}
