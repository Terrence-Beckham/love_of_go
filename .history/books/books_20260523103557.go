package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

var Catalog = map[string]Book{
	"abc": {
		ID:     "abc",
		Title:  "The Mandalorian",
		Author: "George Lucas",
		Copies: 4,
	},
	"def": {
		ID:     "def",
		Title:  "Darth Bane",
		Author: "Alex Karpashy ",
		Copies: 2,
	},
	"jkl": {
		ID:     "jkl",
		Title:  "The Rule of Two",
		Author: "Hurst",
		Copies: 9,
	},
	"mno": {
		ID:     "mno",
		Title:  "Legacy of Evil",
		Author: "Someone",
		Copies: 3,
	},
}

func GetAllBooks(catalog map[string]Book) []Book {
var	newCollection = slices.Collect(maps.Values(catalog))
re
}

func PrintBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}

func GetBook(catalog map[string]Book,ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}
func AddBook(book Book)  {
	Catalog[book.ID] = book
}
