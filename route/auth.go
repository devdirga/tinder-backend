package route

import (
	c "gotinder/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	app.Post("/signup", c.Signup)
	app.Post("/signin", c.Signin)
	app.Get("/verification/:token", c.VerificationToken)
}
