package books

import "fmt"

type Book struct {
	ID    string 
	Title  string
	Author string
	Copies int
}

var catalog = []Book{
	{
		ID:     "abc",
		Title:  "The Mandalorian",
		Author: "George Lucas",
		Copies: 4,
	},
	{
		ID:     "def",
		Title:  "Darth Bane",
		Author: "Alex Karpashy ",
		Copies: 2,
	},
	{
		ID:     "jkl",
		Title:  "The Rule of Two",
		Author: "Hurst",
		Copies: 9,
	},
	{
		ID:     "mno,
		Title:  "Legacy of Evil",
		Author: "Someone",
		Copies: 3,
	},
}

func GetAllBooks() []Book {
	return catalog
}

func PrintBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}

func GetBook(ID int) (Book, bool) {
	for _, book := range catalog {
		if book.ID == ID {
			return book, true
		}
	}

	return Book{}, false
}
