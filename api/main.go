package main

import (
  "github.com/gin-gonic/gin"
  "gotodo/database"
  "gotodo/routes"
)

func main() {

  database.ConnectDatabase()
  // Create a Gin router with default middleware (logger and recovery)
  r := gin.Default()

  // Define a simple GET endpoint
  routes.InitRoutes(r)

  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  r.Run()
}