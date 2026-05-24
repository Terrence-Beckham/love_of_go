package main

import (
	"books"
	"fmt"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))

	for _, book := range books.TestCatalog.GetCatalog() {
		fmt.Println(book.BookToString())
	}

}
