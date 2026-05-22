package books_test 

import (
	"testing"
	"books"

)
func TestBookToString_FormatsBookInfoAsString(t *testing.T){
	input := books.Book{
		Title: "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2 ,
	}

 want := "Sea Room by Adam Nicolson (copies: 2)"
 got := BookToStrin(input)
 
 if want != got {
	t.Fatalf("%q not equal to %q", want, got)
 }
	
}
