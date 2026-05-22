package books

import "fmt"

type Book struct {
	Id int
	Title  string
	Author string
	Copies int
}
var catalog = []Book{
	{
		Id: 1,
		Title: "The Mandalorian",
		Author: "George Lucas",
		Copies: 4,
	},
		{
			Id: 2,
		Title: "Darth Bane",
		Author: "Alex Karpashy ",
		Copies: 2,
	},
		{
			Id:3,
		Title: "The Rule of Two",
		Author: "Hurst",
		Copies: 9,
	},
		{
			Id:4,
		Title: "Legacy of Evil",
		Author: "Someone",
		Copies: 3,
	},
}

func GetAllBooks() []Book {
return catalog	
}

func printBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}

func GetBook(id int, books []Book) Book {
for _, book := range books{
	currentBook  Book{}  book.Id == id
}
}
