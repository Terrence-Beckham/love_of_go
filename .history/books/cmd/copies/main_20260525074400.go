package main

import (
	"books"
	"fmt"
	"os"
	"strconv"
)
func main()  {
if len(os.Args) != 3{
	fmt.Println("Unable to find <Books>")
	return
}
catalog, err	:=books.OpenCatalog("testdata/catalog")
if err != nil{
	fmt.Printf("opening catalog: %v/",err)
	return
}
ID := os.Args[1]
book, ok := catalog.GetBook(ID)
if ! ok {
	fmt.Println("Sorry we couldn't find the book in the catalog.")
	return
}
fmt.Println(book)
updatedCopies := os.Args[2]
newCopies,err :=strconv.Atoi(updatedCopies)
if err != nil{
	fmt.Printf("Incorrect input %v", newCopies)
	return
}
currentCopiescurrentbook.Copies
book.Copies =  newCopies
fmt.Printf("%v had %v copies before. I now has %v copies",book,updatedCopies, newCopies)


}