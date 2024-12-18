package controller

import (
	"errors"
	"gotinder/model"
	"gotinder/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	userData, _ := model.GetUserByEmail(email)
	if user == nil {
		return util.SendRes(c, errors.New("email does not exist"), nil)
	}
	return util.SendRes(c, nil, userData)
}

func UserUpdate(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	var usr model.User
	if err := c.BodyParser(&usr); err != nil {
		return c.Status(400).SendString("Failed to parse request body")
	}
	err := model.UserUpdateByEmail(model.User{
		Email:        email,
		Bio:          usr.Bio,
		ProfileImage: usr.ProfileImage,
	})
	if err != nil {
		return util.SendRes(c, err, nil)
	}
	return util.SendRes(c, nil, nil)
}
