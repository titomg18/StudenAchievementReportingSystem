package routes

import (
    "database/sql"

    repo "StudenAchievementReportingSystem/app/repository/postgresql"
    service "StudenAchievementReportingSystem/app/service/postgresql"
    "StudenAchievementReportingSystem/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, db *sql.DB) {

    // Initialize repository and service
    userRepo := repo.NewUserRepository(db)
    authService := service.NewAuthService(userRepo)

    api := app.Group("/api/v1")

    // Authentication routes
    auth := api.Group("/auth")

    // Public routes
    auth.Post("/login", authService.Login)
    auth.Post("/refresh", authService.Refresh)

    // Protected routes
    auth.Post("/logout", middleware.AuthRequired(), authService.Logout)
    auth.Get("/profile", middleware.AuthRequired(), authService.Profile)
}

     // Students and Lecturers
    auth.Get("/students", studentService.GetAllStudents)
    auth.Get("/students/:id", studentService.GetStudentByID)
    auth.Get("/students/:id/achievements", studentService.GetStudentAchievements)
    auth.Put("/students/:id/advisor", studentService.UpdateAdvisor)
    auth.Get("/lecturers", lecturerService.GetAllLecturers)
    auth.Get("/lecturers/:id/advisees", lecturerService.GetAdvisees)