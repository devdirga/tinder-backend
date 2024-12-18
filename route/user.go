package route

import (
	c "gotinder/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", c.CreateUser)
}
