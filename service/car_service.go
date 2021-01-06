package service

import (
	"gin_test/db"
	"gin_test/entity"

	"github.com/gin-gonic/gin"
)

type CarService struct{}

type Car entity.Car

//車両登録
func (s CarService) CreateModel(c *gin.Context) (Car, error) {
	db := db.GetDB()
	var car Car

	if err := c.BindJSON(&car); err != nil {
		return car, err
	}

	if err := db.Create(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}

//車両全件取得
func (s CarService) GetAll() ([]Car, error) {
	db := db.GetDB()
	var car []Car

	if err := db.Find(&car).Error; err != nil {
		return nil, err
	}

	return car, nil
}

//車両指定件数取得
func (s CarService) GetSum(limit string) ([]Car, error) {
	db := db.GetDB()
	var car []Car

	if err := db.Limit(limit).Find(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}

//IDで指定した車両を取得
func (s CarService) GetByID(carId string) (Car, error) {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).First(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}

//IDで指定した車両を更新
func (s CarService) UpdateByID(carId string, c *gin.Context) (Car, error) {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).First(&car).Error; err != nil {
		return car, err
	}

	if err := c.BindJSON(&car); err != nil {
		return car, err
	}

	db.Save(&car)

	return car, nil
}

//IDで指定した車両を削除
func (s CarService) DeleteByID(carId string) error {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).Delete(&car).Error; err != nil {
		return err
	}

	return nil
}
