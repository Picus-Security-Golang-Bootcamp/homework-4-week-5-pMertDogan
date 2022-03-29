package bookRest

import (
	
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)


/*
ITS CASE SENSIITIVE
PS C:\Projeler\homework-4-week-5-pMertDogan> curl -v "http://localhost:8080/book/find?name=Hob"
VERBOSE: GET with 0-byte payload
VERBOSE: received 369-byte response of content type application/json


StatusCode        : 200
StatusDescription : OK
Content           : {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:00","DeletedAt":null,"ID":"1","AuthorID":"0","BookName":"Hobbit","NumberOfPages":665,"Sto 
                    ckCount":14,"Price":...
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 369
                    Content-Type: application/json
                    Date: Tue, 29 Mar 2022 17:23:15 GMT

                    {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:...
Forms             : {}
Headers           : {[Content-Length, 369], [Content-Type, application/json], [Date, Tue, 29 Mar 2022 17:23:15 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 369



PS C:\Projeler\homework-4-week-5-pMertDogan>
*/
//get book by id
func FindBookByNameWithoutAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	log.Println("Find Requested  name: " + name)

	b, err := book.Repo().FindByName(name)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book is not exist"))
		
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