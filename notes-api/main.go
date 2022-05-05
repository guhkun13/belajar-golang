package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"

	"github.com/guhkun13/belajar-golang/notes-api/database"
	"github.com/guhkun13/belajar-golang/notes-api/router"
)


func main() {
	// start fiber
	fastergoding.Run()
	app := fiber.New()
	
	// connect db
	database.ConnectDB()
	
	// routing
	router.SetupRoutes(app)
	
	// listen
	log.Fatal(app.Listen(":5000"))
}