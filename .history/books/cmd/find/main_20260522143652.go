package main

import (
	"books"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	ID := os.Args[1]
	book,ok := books.GetBook(strconv.FormatInt(ID))
	if !ok{
		fmt.Println("Sorry, I couldn't find that book in the catalog.")
		return
	}
	fmt.Println(books.BookToString(book))
	
}