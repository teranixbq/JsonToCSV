package repository

import (
	"github.com/teranixbq/goJsoncsv/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	Insert(data model.College) error
	Get() ([]model.College, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (eg *repository) Insert(data model.College) error {

	tx := eg.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (eg *repository) Get() ([]model.College, error) {

	dataCollege := []model.College{}

	tx := eg.db.Find(&dataCollege)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return dataCollege, nil

}
