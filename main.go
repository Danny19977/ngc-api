package main

import (
    // "log"
    // "net/http"
    // "github.com/gofiber/fiber/v2"
    // "github.com/Danny19977/ngc-api/tree/main/database"

    // "project/pkg/router"
    // "project/internal/database"
    // "project/internal/logger"
)

// func main() {
//     port := os.Getenv("PORT")
// 	if port == "" {
// 		port = ":8000"
// 	} else {
// 		port = ":" + port
// 	}

// 	return port
// }

// func main() {

// 	database.Connect()

// 	app := fiber.New()

// 	// Initialize default config
// 	app.Use(logger.New())

// 	// Middleware
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins:     "https://mspos.onrender.com, https://mspos-v2.onrender.com, http://localhost:4200",
// 		AllowHeaders:     "Origin, Content-Type, Accept",
// 		AllowCredentials: true,
// 		AllowMethods: strings.Join([]string{
// 			fiber.MethodGet,
// 			fiber.MethodPost,
// 			fiber.MethodHead,
// 			fiber.MethodPut,
// 			fiber.MethodDelete,
// 			fiber.MethodPatch,
// 		}, ","),
// 	}))

// 	routes.Setup(app)

// 	log.Fatal(app.Listen(getPort()))

// }
