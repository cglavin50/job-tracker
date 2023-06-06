package main

import (
	"log"
	"os"

	"github.com/cglavin50/job-tracker/server/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// initialization functions
	database.InitDB(os.Getenv("DSN"))
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000")) //listening on localhost:3000
}
