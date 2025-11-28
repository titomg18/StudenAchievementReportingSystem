package routes

import (
	"github.com/gofiber/fiber/v2"
	"StudenAchievementReportingSystem/app/service"
)

func SetupRoutes(app *fiber.App, authService services.AuthService) {
	AuthRoutes(app, authService)
}
