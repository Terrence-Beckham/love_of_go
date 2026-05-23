package books

import (
	"fmt"
	"maps"
)

type Book struct {
	ID    string 
	Title  string
	Author string
	Copies int
}

var Catalog = map[string]Book{
	"abc":{
		ID:     "abc",
		Title:  "The Mandalorian",
		Author: "George Lucas",
		Copies: 4,
	},
	"def":{
		ID:     "def",
		Title:  "Darth Bane",
		Author: "Alex Karpashy ",
		Copies: 2,
	},
	"jkl":{
		ID:     "jkl",
		Title:  "The Rule of Two",
		Author: "Hurst",
		Copies: 9,
	},
	"mno":{
		ID:     "mno",
		Title:  "Legacy of Evil",
		Author: "Someone",
		Copies: 3,
	},
}

func GetAllBooks() []Book {
	books := maps.Values(Catalog)
	return Catalog
}

func PrintBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}

func GetBook(ID string) (Book, bool) {
	for _, book := range Catalog {
		if book.ID == ID {
			return book, true
		}
	}

	return Book{}, false
}
