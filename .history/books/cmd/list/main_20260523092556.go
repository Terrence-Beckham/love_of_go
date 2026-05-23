package main

import (
	"fmt"
	"books"
)

func main() {
	// fmt.Println(books.BookToString(books.GetAllBooks()))
	var catalog = books.get
	for _,book := range books.GetAllBooks(){
		fmt.Println(books.BookToString(book))
	}

}
