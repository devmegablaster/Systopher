package main

import (
	"github.com/axrav/Systopher/backend/config"
	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Load()
	db.InitPostgres()
	db.InitRedis()
	app := fiber.New()

	app.Use(logger.New())
	// use cors middleware
	app.Use(cors.New(cors.ConfigDefault))
	routes.SetupRoutes(app)

}
