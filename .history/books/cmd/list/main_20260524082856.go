package main

import (
	"books"
	"fmt"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))
catalog := books.GetTestCatalog()
	for _, book := range books.TestCatalog.GetAllBooks() {
		fmt.Println(book.BookToString())
	}

}
