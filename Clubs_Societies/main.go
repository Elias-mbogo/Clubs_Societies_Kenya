package main

import (
	"log"

	routes "exptest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func main() {

	// Create a new engine
	engine := django.New("./views", ".django")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".django"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("base", fiber.Map{
			"Title": "homePage",
		})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("account/signup", fiber.Map{
			"Title": "accountPage",
		})
	})

	app.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("account/signinpath", fiber.Map{
			"Title": "pathPage",
		})
	})

	app.Get("/api/register-club", func(c *fiber.Ctx) error {
		return c.Render("account/registerClub", fiber.Map{
			"Title": "registerClub",
		})
	})

	app.Get("/group", func(c *fiber.Ctx) error {
		return c.Render("users/group", fiber.Map{
			"Title": "groupPage",
		})
	})

	routes.UserRoutes(app)
	routes.ClubRoutes(app)
	app.Static("/", "./public")
	log.Fatal(app.Listen(":3000"))
}
