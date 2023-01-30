package main

import (
	"github.com/codevski/readr/database"
	"github.com/codevski/readr/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
