package bookRest

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
//get book by id
func GetBookByIdIncludeAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookdID := r.URL.Query().Get("bookID")
	// authorInfo := r.URL.Query().Get("authorInfo")
	log.Println("bookID: " + bookdID)
	// log.Println("authorInfo: " + authorInfo)

	b := book.Repo().GetBookByIdIncludeAuthor(bookdID)

	// b := book.Repo().GetByIdWithAuthorName(bookdID)
	// fmt.Println(b)
	if b[0].Name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There is an error occured"))
	}

	w.WriteHeader(http.StatusOK)
	v, err := json.Marshal(b[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Fatal(err)

	}

	w.Write([]byte(v))
}

/*
PS C:\Projeler\homework-4-week-5-pMertDogan> curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
VERBOSE: GET with 0-byte payload
VERBOSE: received 65-byte response of content type application/json


StatusCode        : 200
StatusDescription : OK
Content           : {"ID":1,"BookName":"Hobbit","Name":"J.R.R. Tolkien","AuthorID":0}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 65
                    Content-Type: application/json
                    Date: Tue, 29 Mar 2022 17:08:02 GMT

                    {"ID":1,"BookName":"Hobbit","Name":"J.R.R. Tolkien","AuthorID":0}
Forms             : {}
Headers           : {[Content-Length, 65], [Content-Type, application/json], [Date, Tue, 29 Mar 2022 17:08:02 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 65

*/