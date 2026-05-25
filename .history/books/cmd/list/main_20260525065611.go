package main

import (
	"books"
	"fmt"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))
catalog, err := books.GetCatalog()
	for _, book := range catalog.GetAllBooks() {
		fmt.Println(book.BookToString())
	}

}
