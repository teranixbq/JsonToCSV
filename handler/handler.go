package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teranixbq/goJsoncsv/model"
	"github.com/teranixbq/goJsoncsv/service"
)

type handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *handler {
	return &handler{
		service: service,
	}
}

func (college *handler) Insert(f *fiber.Ctx)error{
	input := model.College{}
	err := f.BodyParser(&input)
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}
	err = college.service.Insert(input)
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": "Failed to insert data",
		})
	}

	return f.Status(201).JSON(fiber.Map{
		"message": "Success insert data",
	})

}

func (college *handler) Get(f *fiber.Ctx)error{

	data, err := college.service.Get()
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": "Failed to get data",
		})
	}

	return f.Status(200).JSON(fiber.Map{
		"message": "Success get data",
		"data": data,
	})

}

func (college *handler) CsvUser(f *fiber.Ctx)error {
	result,err := college.service.GetUserCSV()
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	f.Set(fiber.HeaderContentType, "text/csv")
	f.Set(fiber.HeaderContentDisposition, "attachment; filename=datauser.csv")

	return f.Send(result)
}