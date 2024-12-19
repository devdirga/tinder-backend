package main

import (
	"gotinder/config"
	"gotinder/consumer"
	"gotinder/database"
	"gotinder/middleware"
	"gotinder/model"
	"gotinder/route"
	"log"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
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
	route.ProfileRoute(app)
	route.SwipeRoute(app)

	if config.GetConf().IsQueue {
		// Start Kafka consumer
		go func() {
			if err := consumer.ConsumeMessages(); err != nil {
				log.Fatalf("Error starting consumer: %v", err)
			}
		}()
	}

	log.Fatal(app.Listen(":5000"))
}
