package main

import (
	"log"

	"github.com/vendenta/database"
	"github.com/vendenta/handlers"
)

func main() {
	log.Println("Start Project")

	database.PingDataBase()

	database.AutoMigrateDB()

	handlers.HandRouters()
}
