package main

import (
	"log"

	"github.com/joho/godotenv"
	controller "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller"
	database "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/data/database"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/author"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//init env and parse flags
func init() {
	//Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. \n ERROR : " + err.Error())
	}
	//parse reset flag or any other
	// flag.Parse()

}

func main() {
	//init database
	db , err := database.ConnectPostgresDB()
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	//init book repo
	book.BookRepoInit(db)
	author.AuthorRepoInit(db)
	//setup the endpoint controllers
	controller.SetupControllers()
}
