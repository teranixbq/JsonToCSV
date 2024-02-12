package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teranixbq/goJsoncsv/handler"
	"github.com/teranixbq/goJsoncsv/repository"
	"github.com/teranixbq/goJsoncsv/service"
	"gorm.io/gorm"
)

func RouteInit(f *fiber.App, db *gorm.DB){
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	f.Post("/register", handler.Insert)
	f.Get("/", handler.Get)
	f.Get("/download",handler.CsvUser)
}