package godemo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	handleRequests()
	
	fmt.Println()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home</h1>")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/books", getAllBooks)

	http.ListenAndServe(":8080", nil)
}

type book struct {
	Isbn       string `json:"isbn"`
	Name       string `json:"name"`
	AuthorName string `json:"authorName"`
	Year       string `json:"year"`
	TotalPage  int    `json:"totalPage"`
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	
	allBooks := []book{
		 {
			Isbn:       "978-616-08-2539-4",
			Name:       "พัฒนาเว็บแอปพลิเคชั่นด้วย JavaScript",
			AuthorName: "จตุรพัชร์ พัฒนทรงศิวิไล",
			Year:       "2529",
			TotalPage:  460,
		},
		 {
			Isbn:       "978-616-93108-0-8",
			Name:       "BIG DATA SERIES I, Introduction to a Big Data Project",
			AuthorName: "ดร.อสมา กุลวานิชไชยนันท์",
			Year:       "2561",
			TotalPage:  247,
		},
	}

	json.NewEncoder(w).Encode(allBooks)
}
func whateverPage (w http.ResponseWriter, r *http.Request){

}

//ctrl + c
