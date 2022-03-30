package authorRest

import (
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
)

//curl -v "http://localhost:8080/book?bookID=1"
//get book by id
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// var responseModel domain.APIStructAuthor = domain.APIStructAuthor{}


}
