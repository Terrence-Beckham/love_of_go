package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"sync"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}
type Catalog struct {
	mu *sync.RWMutex
	data map[string]Book
	Path string
}	

func (catalog *Catalog) GetAllBooks() []Book {
	var newCollection = slices.Collect(maps.Values(catalog.data))
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

func (catalog *Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog.data[ID]
	return book, ok
}

func (book *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	fmt.Println("after update book.Copies =", book.Copies)
	return nil
}
func OpenCatalog(path string) (*Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := Catalog{
		mu: &sync.RWMutex{},
		data: map[string]Book{},
	}
	err = json.NewDecoder(file).Decode(&catalog.data)
	if err != nil {
		return nil, err
	}
	catalog.Path = path
	return &catalog, nil
}

func (catalog *Catalog) Sync(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(catalog)
	if err != nil {
		return err
	}
	return nil
}
func (catalog *Catalog) SetCopies(ID string, copies int) error {
	book, ok := catalog.data[ID]
	if !ok {
		return fmt.Errorf("ID %q not found", ID)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	catalog.data[ID] = book
	return nil

}
func (catalog *Catalog) AddBook(book Book) error {
	_, ok := catalog.data[book.ID]
	if ok {
		return fmt.Errorf("ID %q already exists", book.ID)
	}
	catalog.data[book.ID] = book
	return nil
}
func (catalog *Catalog) GetCopies(ID string) (int, error) {
	book, ok := catalog.data[ID]
	if !ok {
		return 0, fmt.Errorf("ID %q is not found", ID)

	}
	return book.Copies, nil
}
