package routes

// import (
// 	service "StudenAchievementReportingSystem/app/Service"

// 	"github.com/gofiber/fiber/v2"
// )

// func SetupRoutes(app *fiber.App) {
// 	// Auth routes
// 	auth := app.Group("/api/auth")
// 	auth.Post("/login", service.Claims)
	
// 	// User routes
// 	users := app.Group("/api/users")
// 	users.Get("/")
// 	users.Get("/:id")
// 	users.Post("/")
// 	users.Put("/:id")
// 	users.Delete("/:id")

// 	// // Student achievement routes
// 	// achievements := app.Group("/api/achievements")
// 	// achievements.Get("/")
// 	// achievements.Get("/:id")
// 	// achievements.Post("/")
// 	// achievements.Put("/:id")
// 	// achievements.Delete("/:id")

// 	// // Report routes
// 	// reports := app.Group("/api/reports")
// 	// reports.Get("/")
// 	// reports.Get("/:id")
// 	// reports.Post("/")
// 	// reports.Get("/export")

// 	// // Health check
// 	// app.Get("/health")
// }