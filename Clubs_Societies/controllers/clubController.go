package controllers

import (
	db "exptest/database"
	"exptest/models"

	"github.com/gofiber/fiber/v2"
)

type RegisterClub struct {
	Club  string `json:"club"`
	Group string `json:"group"`
}

func CreateClub(c *fiber.Ctx) error {
	register := RegisterClub{}

	err := c.BodyParser(&register)

	if err != nil {
		c.Status(422).JSON(
			&fiber.Map{"message": "Please input the details correctly"})
		return err
	}

	club := models.Clubs{Name: register.Club}
	err = db.DB.Create(&club).Error
	if err != nil {
		c.Status(422).JSON(
			&fiber.Map{"message": "could not create club"})
		return err
	}

	group := models.Groups{Name: register.Group, ClubsID: club.ID}
	err = db.DB.Create(&group).Error
	if err != nil {
		c.Status(422).JSON(
			&fiber.Map{"message": "could not create group"})
		return err
	}

	c.Status(201).JSON(
		&fiber.Map{"message": "created the club successfully"})
	return c.Redirect("/")

}
