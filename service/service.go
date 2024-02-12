package service

import (
	"encoding/csv"
	"os"

	"github.com/teranixbq/goJsoncsv/model"
	"github.com/teranixbq/goJsoncsv/repository"
)

type service struct {
	repository repository.RepositoryInterface
}

type ServiceInterface interface {
	Insert(data model.College) error
	Get() ([]model.College, error)
	GetUserCSV() error
}

func NewService(repository repository.RepositoryInterface) ServiceInterface {
	return &service{
		repository: repository,
	}
}

func (college *service) Insert(data model.College) error {

	err := college.repository.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

func (college *service) Get() ([]model.College, error) {

	dataCollege, err := college.repository.Get()
	if err != nil {
		return nil, err
	}

	return dataCollege, nil
}

func (college *service) GetUserCSV() error {
	data, err := college.Get()
	if err != nil {
		return err
	}

	outputFile, err := os.Create("public/datauser.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"Nim", "Name", "Campus"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range data {
		var datacsv []string
		datacsv = append(datacsv, r.Nim, r.Name, r.Campus)
		if err := writer.Write(datacsv); err != nil {
			return err
		}
	}

	return nil
}
