package route

import (
	c "gotinder/controller"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(app *fiber.App) {
	app.Get("/me", c.Me)
	app.Post("/me", c.UserUpdate)
}
