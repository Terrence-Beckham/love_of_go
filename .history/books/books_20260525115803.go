package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}
type Catalog map[string]Book



func (catalog Catalog) GetAllBooks() []Book {
	var newCollection = slices.Collect(maps.Values(catalog))
	for _, book := range newCollection {
		fmt.Printf("This is the book %#v", book)
	}
	return newCollection
}

func PrintBook(book Book) {
	fmt.Printf("%v by %v - %v copies", book.Title, book.Author, book.Copies)

}
func (book Book) BookToString() string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)

}

func (catalog Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}
func (catalog Catalog) AddBook(book Book) {
	catalog[book.ID] = book
}
func (book *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	fmt.Println("after update book.Copies =", book.Copies)
	return nil
}
func OpenCatalog(path string) (Catalog, error){
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := Catalog{}
	var decoder = json.NewDecoder(file)
	err = decoder.Decode(&catalog)
	// err = json.NewDecoder(file).Decode(&catalog)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}


func ()  {
	
}