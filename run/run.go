package run

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/ryszhio/goauth/database"
	"github.com/ryszhio/goauth/internal/generator"
	"github.com/ryszhio/goauth/internal/keys"
	"github.com/ryszhio/goauth/router"
)

func InitializeApp() error {
	// Load Environment Files
	godotenv.Load(".env")
	// Generate Keys
	err := keys.ConfigureKeys()
	if err != nil {
		return err
	}
	// Initialize Node number for the snowflake generator
	generator.InitializeNode()
	// Establish connection to our database.
	database.ConnectDB()
	// Initialize fiber app
	app := fiber.New()

	// Setup routes
	router.SetupRoutes(app)

	// Handle root routes
	app.Get("/", greetingResponse)
	app.Post("/", greetingResponse)

	// Serve the application
	app.Listen(":"+os.Getenv("PORT"), fiber.ListenConfig{EnablePrefork: false})

	return nil
}

func greetingResponse(c fiber.Ctx) error {
	return c.SendString("goauth is running\nMade with ❤️ by ryszhio.")
}
