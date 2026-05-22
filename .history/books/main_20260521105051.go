package main
import "fmt"

type Book struct{
	Title string
	Author string
	Copies int
}

func printBook(book Book)  {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)
	
}
func BookToString(book Book) book  {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)
	
}

func main()  {
	fmt.Println("This is it ")
	
}