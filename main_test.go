package main

import (
	"bytes"
	"encoding/json"
	"gotinder/config"
	"gotinder/database"
	"gotinder/middleware"
	"gotinder/model"
	"gotinder/route"
	"net/http"
	"net/http/httptest"
	"testing"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	config.Init()
	database.Migrate()
	model.InitDB()
	app := fiber.New()
	app.Use(middleware.CheckSwipeLimit)
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	route.AuthRoute(app)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetConf().Secret)},
	}))
	route.UserRoute(app)
	route.SwipeRoute(app)
	return app
}

func TestAuthSignup(t *testing.T) {
	app := setupTestApp()
	user := model.User{
		Username:     "ahmad20",
		Email:        "ahmad20@gmail.com",
		Password:     "ahmad20",
		Bio:          "bio",
		ProfileImage: "Profile",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthSignin(t *testing.T) {
	app := setupTestApp()
	user := model.User{
		Email:    "ahmad20@gmail.com",
		Password: "ahmad20",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetSwipe(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/swipe", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhbnRvcm8uZGV2ZWxvcGVyQGdtYWlsLmNvbSIsImV4cCI6MTc2NTU4MTM0NCwiaWQiOiI5MWY2ZTI5Yy1iMjE5LTQ5NTMtYTIxNi1hNjk2NTk4OWMwYWUiLCJ1c2VybmFtZSI6ImRpcmdhIn0.7j-86pdaqufmdDba4RpG6k9PIe1E58gdlxblUGjW-PI")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSwipe(t *testing.T) {
	app := setupTestApp()
	parsedUUID, _ := uuid.Parse("3b438420-d3cc-4420-b026-1dae01488489")
	user := model.Swipe{
		TargetUserID: parsedUUID,
		SwipeType:    "like",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/swipe", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhbnRvcm8uZGV2ZWxvcGVyQGdtYWlsLmNvbSIsImV4cCI6MTc2NTU4MTM0NCwiaWQiOiI5MWY2ZTI5Yy1iMjE5LTQ5NTMtYTIxNi1hNjk2NTk4OWMwYWUiLCJ1c2VybmFtZSI6ImRpcmdhIn0.7j-86pdaqufmdDba4RpG6k9PIe1E58gdlxblUGjW-PI")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
