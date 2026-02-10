package main

import (
	"log"

	"github.com/vKousik/go-gin-inventory/database"
	"github.com/vKousik/go-gin-inventory/internal/app"
)

func main() {
	db := database.ConnectDB()
	r := app.NewRouter(db)

	log.Println("server running on :8080")
	r.Run(":8080")
}
