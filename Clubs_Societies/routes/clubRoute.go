package routes

import (
	"exptest/controllers"

	"github.com/gofiber/fiber/v2"
)

func ClubRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register-club", controllers.CreateClub)
}
