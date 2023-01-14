package main

import (
	"CRUDTest/internal/book"
	"CRUDTest/internal/db"
	"CRUDTest/internal/service"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	router := service.NewServer()

	db, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	handler := book.NewHandler(db)
	handler.Register(router.Router)

	router.Run()
}
