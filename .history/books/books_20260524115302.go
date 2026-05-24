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
type Catalog map[string]Book

func GetCatalog() Catalog {
	return Catalog{
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

func (catalog Catalog) GetAllBooks() []Book {
	var newCollection = slices.Collect(maps.Values(catalog))
	for _, book := range newCollection {
		fmt.Printf("This is the book %#v", book)
	}
	return newCollection
}

func PrintBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func (book Book) BookToString() string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}

func (catalog Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}
func (catalog Catalog) AddBook(book Book) {
	catalog[book.ID] = book
}
func (book* Book) SetCopies(copies int) Book {
	fmt.Println("before update book.Copies =",book.Copies)
	book.Copies = copies
	fmt.Println("after update book.Copies =",book.Copies)
	return book
}
