package main

import (
    "log"
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/Danny19977/ngc-api/database"

    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
)

func getPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000" // Default port if not specified in .env
    }
    return ":" + port
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Connect to the database
    if err := database.Connect(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    app := fiber.New()

    // Initialize default config
    app.Use(logger.New())

    // Middleware
    app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin, Content-Type, Accept",
        AllowCredentials: true,
        AllowMethods: strings.Join([]string{
            fiber.MethodGet,
            fiber.MethodPost,
            fiber.MethodHead,
            fiber.MethodPut,
            fiber.MethodDelete,
            fiber.MethodPatch,
        }, ","),
    }))

    // Setup routes (if any)
    // routes.Setup(app)

    // Start the server
    port := getPort()
    log.Printf("Server is running on port %s", port)
    if err := app.Listen(port); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
