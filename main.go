package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"StudenAchievementReportingSystem/routes"
	"StudenAchievementReportingSystem/database"
	"StudenAchievementReportingSystem/app/repository"
	"StudenAchievementReportingSystem/app/service"
)

func main() {
	app := fiber.New()

	db := database.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, "JWT_SECRET_123")

	routes.SetupRoutes(app, authService)

	log.Println("Server running on :3000")
	app.Listen(":3000")
}
