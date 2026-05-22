package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func ()  {
	
}

func printBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}
