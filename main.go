package main

import (
	"log"
	"StudenAchievementReportingSystem/routes"  // ← PERBAIKI IMPORT PATH

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Println("Server starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}