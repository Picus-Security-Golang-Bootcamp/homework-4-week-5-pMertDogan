package bookRest

import (
	
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//curl -v "http://localhost:8080/book?bookID=1"
//get book by id
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookdID := r.URL.Query().Get("bookID")
	authorInfo := r.URL.Query().Get("authorInfo")
	log.Println("bookID: " + bookdID)
	log.Println("authorInfo: " + authorInfo)

	b, err := book.Repo().GetByID(bookdID)
	// b := book.Repo().GetByIdWithAuthorName(bookdID)
	// fmt.Println(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book is not exist"))
		//Good idea? exit?
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	v, err := b.Marshal()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Fatal(err)

	}

	w.Write([]byte(v))
}


/*

PS C:\Projeler\homework-4-week-5-pMertDogan> curl -v "http://localhost:8080/book?bookID=1"
VERBOSE: GET with 0-byte payload
VERBOSE: received 369-byte response of content type application/json


StatusCode        : 200
StatusDescription : OK
Content           : {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:00","DeletedAt":null,"ID":"1","AuthorID":"0","BookName":"Hobbit","NumberOfPages":665,"Sto 
                    ckCount":14,"Price":...
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 369
                    Content-Type: application/json
                    Date: Tue, 29 Mar 2022 17:06:58 GMT

                    {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:...
Forms             : {}
Headers           : {[Content-Length, 369], [Content-Type, application/json], [Date, Tue, 29 Mar 2022 17:06:58 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 369

*/