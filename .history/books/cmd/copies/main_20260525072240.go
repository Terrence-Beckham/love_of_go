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
func GetCatalog() Catalog {
	return Catalog{
		"def": {
			ID:     "def",
			Title:  "Darth Bane",
			Author: "Alex Karpashy ",
			Copies: 2,
		},
		"abc": {
			ID:     "abc",
			Title:  "The Mandalorian",
			Author: "George Lucas",
			Copies: 4,
		},
	}
}
}