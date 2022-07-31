package main

import (
	"os"
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", routes.ResolveURL) 
	app.Post("/", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}