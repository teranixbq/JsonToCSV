package service

import (
	"github.com/teranixbq/goJsoncsv/model"
	"github.com/teranixbq/goJsoncsv/repository"
)

type service struct {
	repository repository.RepositoryInterface
}


type ServiceInterface interface {
	Insert(data model.College) error
	Get() ([]model.College, error)
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
