package controller

import (
	"errors"
	"gotinder/config"
	"gotinder/model"
	"gotinder/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	payload := new(model.User)
	if err := c.BodyParser(payload); err != nil {
		return util.SendRes(c, err, nil)
	}

	if user, _ := model.GetUserByEmail(payload.Email); user != nil {
		return util.SendRes(c, errors.New("email already in use"), nil)
	}

	if bytes, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password+config.GetConf().Secret),
		bcrypt.DefaultCost,
	); err != nil {
		return util.SendRes(c, err, nil)
	} else {
		if err = model.UserCreate(model.User{
			Username:         payload.Username,
			Email:            payload.Email,
			Password:         string(bytes),
			Bio:              payload.Bio,
			ProfileImage:     payload.ProfileImage,
			SubscriptionType: "free",
		}); err != nil {
			return util.SendRes(c, err, nil)
		} else {
			model.VerifTokenCreate(model.VerifToken{Email: payload.Email})
			return util.SendRes(c, err, nil)
		}
	}
}

func VerificationToken(c *fiber.Ctx) error {
	return util.SendRes(c, model.VerifTokenConfirm(c.Params("token")), nil)
}

func Signin(c *fiber.Ctx) error {
	payload := new(model.User)
	if err := c.BodyParser(payload); err != nil {
		return util.SendRes(c, err, nil)
	}
	user, _ := model.GetUserByEmail(payload.Email)
	if user == nil {
		return util.SendRes(c, errors.New("email does not exist"), nil)
	}
	if ok, _ := util.CompareHash(
		user.Password, payload.Password, config.GetConf().Secret,
	); !ok {
		return util.SendRes(c, errors.New("password incorrect"), nil)
	}
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      util.GetNow().Add(time.Hour * 8640).Unix(), // 1 year
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GetConf().Secret))
	return util.SendRes(c, err, fiber.Map{
		"token":    t,
		"username": user.Username,
	})
}
