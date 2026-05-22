package main
import "fmt"



func printBook(book Book)  {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)
	
}
func BookToString(book Book) string  {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	
}

func main()  {
	book := Book{
		Title: "Engineering in Plain Sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}
	fmt.Println(BookToString(book))
	
}