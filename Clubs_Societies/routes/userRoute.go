package routes

import (
	"exptest/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/signup", controllers.CreateUser)
}
