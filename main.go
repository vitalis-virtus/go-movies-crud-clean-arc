package main

import (
	"fmt"
	"github.com/vitalis-virtus/go-movies-gallery/pkg/database"
	"github.com/vitalis-virtus/go-movies-gallery/server"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	app := server.NewApp(db)

	fmt.Println("Starting server on port 8080")

	if err := app.Run("8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
