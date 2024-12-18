package controller

import (
	"gotinder/model"
	"gotinder/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func SwipeCreate(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	var swipe model.Swipe
	if err := c.BodyParser(&swipe); err != nil {
		return c.Status(400).SendString("Failed to parse request body")
	}

	parsedUUID, _ := uuid.Parse(id)
	swipe.UserID = parsedUUID
	if err := model.SwipeCreate(swipe); err != nil {
		return c.Status(500).SendString("Failed to create swape:" + err.Error())
	}
	return c.Status(201).JSON(swipe)
}

func SwipeData(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	data, err := model.SwipeData(id)
	if err != nil {
		return c.Status(400).SendString("Failed to parse request body")
	}

	if len(data) > 0 {
		// for _, v := range data {
		// 	if v.Email == "ahmad21@gmail.com" {
		// 		return util.SendRes(c, err, v)
		// 	}
		// }
		return util.SendRes(c, err, data[0])
	}

	return util.SendRes(c, err, data)
}
