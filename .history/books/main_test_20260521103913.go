package main

import "testing"
func TestBookToString_FormatsBookInfoAsString(t *testing.T){
	input := Book{
		Title: "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2 ,
	}

 want := "Sea Room by Adam Nicolson - 2 copies"
 got := TestBookToString()
 
 if want != got {
	t.Fatalf("")
 }
	
}
}