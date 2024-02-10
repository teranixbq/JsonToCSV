package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teranixbq/goJsoncsv/config"
	routes "github.com/teranixbq/goJsoncsv/route"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main(){
	f := fiber.New()
	
	db := config.InitPostgresDB()
	routes.RouteInit(f, db)

	f.Use(cors.New())

	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	f.Listen(":3000")
}