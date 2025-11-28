package routes

import (
	"github.com/gofiber/fiber/v2"
	"StudenAchievementReportingSystem/app/models"
	"StudenAchievementReportingSystem/app/service"
)

func AuthRoutes(app *fiber.App, authService services.AuthService) {

	app.Post("/login", func(c *fiber.Ctx) error {

		var req models.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid request",
			})
		}

		// panggil service (logic di service, bukan di route)
		resp, err := authService.Login(&req)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// success
		return c.JSON(resp)
	})

		app.Post("/register", func(c *fiber.Ctx) error {
		var req models.RegisterRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		// panggil service
		if err := authService.Register(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "User registered successfully",
		})
	})

}

