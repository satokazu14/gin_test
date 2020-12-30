package car

import (
	"gin_test/db"
	"gin_test/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

type Car entity.Car

//車両登録
func (s Service) CreateModel(c *gin.Context) (Car, error) {
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
func (s Service) GetAll() ([]Car, error) {
	db := db.GetDB()
	var car []Car

	if err := db.Find(&car).Error; err != nil {
		return nil, err
	}

	return car, nil
}

//車両指定件数取得
func (s Service) GetSum(limit string) (Car, error) {
	db := db.GetDB()
	var car []Car

	if err := db.Where("limit = ?", limit).(car).Error; err != nil {
		return car, err
	}

	return car, nil
}

//IDで指定した車両を取得
func (s Service) GetByID(carId string) (Car, error) {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).First(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}

//IDで指定した車両を更新
func (s Service) UpdateByID(carId string, c *gin.Context) (Car, error) {
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
func (s Service) DeleteByID(carId string) error {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).Delete(&car).Error; err != nil {
		return err
	}

	return nil
}
