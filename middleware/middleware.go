package middleware

import (
	"gotinder/config"
	"gotinder/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckSwipeLimit(c *fiber.Ctx) error {
	if c.Path() == "/swipe" && c.Method() == fiber.MethodPost {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token required",
			})
		}
		// parse
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}
		// Parse JWT token (replace "your-secret-key" with your actual secret)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConf().Secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}
		// Extract user ID from token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["id"] == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}
		email := claims["email"].(string)
		user, err := model.GetUserByEmail(email)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		if user.SubscriptionType == "free" && user.SwipeCount >= 10 {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Swipe limit reached for free subscription",
			})
		}

	}
	return c.Next()
}
