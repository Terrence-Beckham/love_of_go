package main

import(
	"books"
	"fmt"
)

func main()  {
	boo,ok := books.GetBook(2)
	if !ok{
		fmt.Println("Sorry, I couldn't find that book in the catalog.")
		return
	}
	fm
	
}