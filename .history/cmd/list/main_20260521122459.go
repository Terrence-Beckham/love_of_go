package main
import {"fmt"





func main()  {
	book := Book{
		Title: "Engineering in Plain Sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}
	fmt.Println(BookToString(book))
	
}