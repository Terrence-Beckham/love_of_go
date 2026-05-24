package main

import (
	"fmt"
	"books"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))
	
	for _,book := range books.LocalCatalog.GetAllBooks(){
		fmt.Println(books.BookToString(book))
	}

}
