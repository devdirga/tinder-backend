package controller

import (
	"gotinder/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Failed to parse request body")
	}
	if err := model.UserCreate(user); err != nil {
		return c.Status(500).SendString("Failed to create user:" + err.Error())
	}
	return c.Status(201).JSON(user)
}
