package main

import (
	"fmt"
	"github.com/vitalis-virtus/go-movies-gallery/api"
	"github.com/vitalis-virtus/go-movies-gallery/pkg/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	router := api.SetupRouter(db)
	http.Handle("/", router)
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
