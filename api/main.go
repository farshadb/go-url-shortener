package main

import (
	"os"
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	//"github.com/F4r5h4d/go-url-shortener/api"
	"go-url-shortener/api/routes"
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