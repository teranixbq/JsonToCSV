package service

import (
	"bytes"
	"encoding/csv"

	"github.com/teranixbq/goJsoncsv/model"
	"github.com/teranixbq/goJsoncsv/repository"
)

type service struct {
	repository repository.RepositoryInterface
}

type ServiceInterface interface {
	Insert(data model.College) error
	Get() ([]model.College, error)
	GetUserCSV() ([]byte, error)
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

func (college *service) GetUserCSV() ([]byte, error) {
	data, err := college.Get()
	if err != nil {
		return nil, err
	}

	csvbuffer:= bytes.Buffer{}
	writer := csv.NewWriter(&csvbuffer)

	header := []string{"Nim", "Name", "Campus"}
	if err := writer.Write(header); err != nil {
		return nil, err
	}

	for _, v := range data {
		var datacsv []string
		datacsv = append(datacsv, v.Nim, v.Name, v.Campus)
		if err := writer.Write(datacsv); err != nil {
			return nil, err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	return csvbuffer.Bytes(), nil
}