package controllers

import (
	"encoding/json"
	
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//curl -v "localhost:8080/book?bookID=1"
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
