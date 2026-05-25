package main

import (
	"books"
	"fmt"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))
catalog, err := books.OpenCatalog("testdata/")
	for _, book := range catalog.GetAllBooks() {
		fmt.Println(book.BookToString())
	}

}
