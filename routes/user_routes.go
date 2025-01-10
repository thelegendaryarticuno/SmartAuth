package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-app/controllers"
)

func UserRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
