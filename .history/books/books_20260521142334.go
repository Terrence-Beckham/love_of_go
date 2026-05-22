package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}
var catalog = []Book{
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
		Title: "The Mandalorian",
		Author: "George Lucas",
		Copies: 4,
	},
}

func GetAllBooks() []Book {
return []Book{}	
}

func printBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}
