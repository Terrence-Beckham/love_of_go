package main
import(
	"books"
	"fmt"
	"os"
)
func main()  {
if len(os.Args) != 2{
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
}