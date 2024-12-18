package route

import (
	c "gotinder/controller"

	"github.com/gofiber/fiber/v2"
)

func SwipeRoute(app *fiber.App) {
	app.Post("/swipe", c.SwipeCreate)
	app.Get("/swipe", c.SwipeData)
}
