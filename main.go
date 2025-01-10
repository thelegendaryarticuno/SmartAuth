package main

import (
	"log"

	"gofiber-app/database"
	"gofiber-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the database
	database.ConnectMongo()

	// Create a Fiber app
	app := fiber.New()

	// Static files for CSS/JS
	app.Static("/static", "./static")
	app.Static("/js", "./js")
	app.Static("/img", "./img")

	// Serve templates for HTML
	app.Static("/", "./templates")

	// Serve index.html as root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/index.html")
	})

	// Routes
	routes.UserRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
