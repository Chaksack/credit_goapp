package main

import (
	"log"
	"xactscore/database"
	"xactscore/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	routes.Setup(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Welcome to xactscore")
	// })

	log.Fatal(app.Listen(":3000"))
}
