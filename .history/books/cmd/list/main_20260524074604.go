package main

import (
	"fmt"
	"books"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))
	
	for _,book := range books.Catalog{}GetAllBooks(books.Catalog){
		fmt.Println(books.BookToString(book))
	}

}
