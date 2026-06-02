package books_test

import (
	"books"
	"cmp"
	"encoding/json"
	"net"
	"net/http"
	"slices"
	"testing"
)

func getTestCatalog() *books.Catalog {
	catalog := books.NewCatalog()
	err := catalog.AddBook(books.Book{
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		ID:     "abc",
	})
	if err != nil {
		panic(err)
	}
	err = catalog.AddBook(books.Book{
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		ID:     "xyz",
	})
	if err != nil {
		panic(err)
	}
	return catalog
}

func assertTestBooks(t *testing.T, got []books.Book) {
	t.Helper()
	want := []books.Book{
		{
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID:     "abc",
		},
		{
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID:     "xyz",
		},
	}
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v,got %#v", want, got)
	}

}
func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	var catalog = getTestCatalog()
	bookList := catalog.GetAllBooks()
	assertTestBooks(t, bookList)
}
func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		ID:     "abc",
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := input.String()
	if want != got {
		t.Fatalf("%q not equal to %q", want, got)
	}
}

func TestGetBook_ReturnCorrectBook(t *testing.T) {
	t.Parallel()

	var catalog = getTestCatalog()
	want := books.Book{

		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}
	got, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("Book not found")
	}

	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}

}
func TestGetBook_ReturnsFalseIfNotFound(t *testing.T) {
	t.Parallel()
	var catalog = getTestCatalog()
	_, ok := catalog.GetBook("sdfs")
	if ok {
		t.Fatal("Nonexistent ID")
	}

	if ok {
		t.Fatal("Want false for nonexistent ID got true")
	}

}

func TestAddBook(t *testing.T) {
	t.Parallel()

	var catalog = getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("The book already exists")
	}
	catalog.AddBook(books.Book{
		ID:     "123",
		Title:  "New Glory",
		Author: "Me Again",
		Copies: 25,
	})

	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}
func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T) {
	t.Parallel()
	book := books.Book{
		Copies: 5,
	}
	err := book.SetCopies(12)
	if err != nil {
		t.Fatal(err)
	}

	if book.Copies != 12 {
		t.Errorf("want 12 copies, got %d", book.Copies)
	}
}
func TestSetCopies_ReturnsErrorIfCopiesNegative(t *testing.T) {
	t.Parallel()
	book := books.Book{}
	err := book.SetCopies(-1)
	if err == nil {
		t.Error("want error for negative copies, got nil")
	}
}

// func TestOpenCatalog_LoadsCatalogDataFromFile(t *testing.T) {
// 	t.Parallel()
// 	catalog, err := books.OpenCatalog("testdata/catalog")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	bookList := catalog.GetAllBooks()
// 	assertTestBooks(t, bookList)
// }

func TestOpenCatalog_ReadsSameDataWrittenBySync(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	catalog.Path = t.TempDir() + "/catalog"
	err := catalog.Sync()
	if err != nil {
		t.Fatal(err)
	}
	newCatalog, err := books.OpenCatalog(catalog.Path)
	if err != nil {
		t.Fatal(err)
	}
	booklist := newCatalog.GetAllBooks()
	assertTestBooks(t, booklist)

}

func TestSetCopies_OnCatalogModifiesSecifiedBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	book, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not found")
	}
	if book.Copies != 1 {
		t.Fatalf("want 1 copy before change, got %d", book.Copies)
	}
	err := catalog.SetCopies("abc", 2)
	if err != nil {
		t.Fatal(err)
	}
	book, ok = catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not found")
	}
	if book.Copies != 2 {
		t.Fatalf("want 2 copies after change, got %d", book.Copies)

	}
}
func TestAddBook_ReturnsErrorIfIDExists(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not present")
	}
	err := catalog.AddBook(books.Book{
		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	})
	if err == nil {
		t.Fatal("want error for duplicate ID, got nil")
	}
}
func TestAddBook_AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("book already present")
	}
	err := catalog.AddBook(books.Book{
		ID:     "123",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}

func TestSetCopies_IsRaceFree(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	go func() {
		for range 100 {
			err := catalog.SetCopies("abc", 0)
			if err != nil {
				panic(err)
			}
		}
	}()
	for range 100 {
		_, err := catalog.GetCopies("abc")
		if err != nil {
			t.Fatal(err)
		}
	}

}
func TestNewCatalog_CreateEmptyCatalog(t *testing.T) {
	t.Parallel()
	catalog := books.NewCatalog()
	books := catalog.GetAllBooks()

	if len(books) > 0 {
		t.Errorf("want empty catalog, got %#v", books)

	}
}


func TestServerListsAllBooks(t *testing.T){
	t.Parallel()
	catalog := getTestCatalog()
	catalog.Path = t.TempDir() + "/catalog"
	go func(){
		err:= books.ListenAndServe(":3000",catalog)
		if err != nil {
			panic(err)
		}
	}()
	resp, err :=http.Get("http://localhost:3000")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		t.Fatalf("unexpected status %d",resp.StatusCode)
	}
	bookList := []books.Book{}
		err = json.NewDecoder(resp.Body).Decode(&bookList)
		if err != nil {
			t.Fatal(err)
		}
		assertTestBooks(t,bookList)
	
	}

func randomLocalAddr(t *testing.T)string  {
	t.Helper()
	l, err := net.Listen("tcp",":0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return l.Addr().String()
}