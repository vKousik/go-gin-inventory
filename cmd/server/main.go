package main

import (
	"log"

	"github.com/vKousik/go-gin-inventory/internal/app"
	"github.com/vKousik/go-gin-inventory/internal/infrastructure/database"
)

func main() {
	db := database.ConnectDB()
	r := app.NewRouter(db)

	log.Println("server running on :8080")
	r.Run(":8080")
}
