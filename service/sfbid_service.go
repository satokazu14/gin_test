package service

import (
	"gin_test/db"
	"gin_test/entity"
)

type SfbidService struct{}

type SfbidCar entity.SfbidCar

func (ss SfbidService) GetAll() ([]SfbidCar, error) {
	db := db.GetDB()
	var sf []SfbidCar

	if err := db.Find(&sf).Error; err != nil {
		return nil, err
	}

	return sf, nil
}
