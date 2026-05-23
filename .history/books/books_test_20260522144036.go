package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	want := []books.Book{
		{ID: "abc",
			Title:  "The Mandalorian",
			Author: "George Lucas",
			Copies: 4,
		},
		{ID: "def",
			Title:  "Darth Bane",
			Author: "Alex Karpashy ",
			Copies: 2,
		},
		{ID: "ghi",
			Title:  "The Rule of Two",
			Author: "Hurst",
			Copies: 9,
		},
		{ID: "jkl",
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
		ID:     "abc",
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
		ID:     "lmn",
		Title:  "Legacy of Evil",
		Author: "Someone",
		Copies: 3,
	}
	got, ok := books.GetBook(4)
	if !ok {
		t.Fatal("Book not found")
	}

	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}
func TestGetBook_ReturnsFalseIfNotFound(t *testing.T) {
	t.Parallel()

	_, ok := books.GetBook("sdfs")
	if ok {
		t.Fatal("Nonexistent ID")
	}

	if ok {
		t.Fatal("Want false for nonexistent ID got true")
	}

}
