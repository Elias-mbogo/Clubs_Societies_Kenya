package controllers

import (
	db "exptest/database"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"passwords"`
}

func CreateUser(c *fiber.Ctx) error {
	user := User{}

	err := c.BodyParser(&user)

	if err != nil {
		c.Status(422).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = db.DB.Create(&user).Error
	if err != nil {
		c.Status(422).JSON(
			fiber.Map{"message": "could not create user"})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"message": "user has been added"})
	return c.Redirect("/signin")

}
