package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gofiber/fiber/v2"

    "StudenAchievementReportingSystem/config"
    "StudenAchievementReportingSystem/database"
    "StudenAchievementReportingSystem/routes"
)

func main() {

    // 1. Load .env
    config.LoadEnv()

    // 2. Connect PostgreSQL
    database.ConnectPostgres()
    defer database.PostgresDB.Close()

    // 3. Connect MongoDB (jika ada)
    database.ConnectMongo()

    // 4. Inisialisasi Fiber
    app := fiber.New()

    // 5. Setup Routes
    routes.SetupAuthRoutes(app, database.PostgresDB)

    fmt.Println("Setup route berhasil")

    // 6. Start Server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server running on :%s", port)

    if err := app.Listen(":" + port); err != nil {
        log.Fatalf("Server stopped: %v", err)
    }
}
