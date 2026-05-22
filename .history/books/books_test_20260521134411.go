package books_test 

import (
	"testing"
	"books"

)
func TestGetAllBooks(t *testing)  {
want := []books.Book{
	{Title: "in the Company of Cheeful Ladies",
	Author: "Alexander Mc",
}

}
func TestBookToString_FormatsBookInfoAsString(t *testing.T){
	input := books.Book{
		Title: "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2 ,
	}

 want := "Sea Room by Adam Nicolson (copies: 2)"
 got := books.BookToString(input)
 
 if want != got {
	t.Fatalf("%q not equal to %q", want, got)
 }
	
}
