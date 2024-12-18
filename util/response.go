package util

import "github.com/gofiber/fiber/v2"

func SendRes(c *fiber.Ctx, message error, data interface{}) error {
	if message == nil {
		return c.JSON(fiber.Map{"error": false, "message": "success", "data": data})
	} else {
		return c.JSON(fiber.Map{"error": true, "message": message.Error(), "data": data})
	}
}
