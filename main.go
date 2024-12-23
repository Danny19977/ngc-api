package main

import (
    "log"
    "net/http"
    "project/pkg/router"
    "project/internal/database"
    "project/internal/logger"
)

func main() {
    logger.Setup()
    db, err := database.Initialize()
    if err != nil {
        log.Fatalf("Could not set up database: %v", err)
    }
    defer db.Close()

    r := router.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080", r))
}
