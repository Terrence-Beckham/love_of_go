package main

import (
	"books"
	"fmt"
	"os"
)

func main()  {
	ID := os.Args
	book,ok := books.GetBook(2)
	if !ok{
		fmt.Println("Sorry, I couldn't find that book in the catalog.")
		return
	}
	fmt.Println(books.BookToString(book))
	
}