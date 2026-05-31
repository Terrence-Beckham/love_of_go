package main

import (
	"books"
	"fmt"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))
catalog, err := books.OpenCatalog("testdata/catalog")
if err != nil{
		fmt.Printf("opening catalog: %v/n", err)
		return
}
	for _, book := range catalog.GetAllBooks() {
		fmt.Println(book.String())
	}

}
