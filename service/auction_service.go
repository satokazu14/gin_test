package service

import (
	"gin_test/db"
	"gin_test/entity"

	"github.com/gin-gonic/gin"
)

type AuctionService struct{}

type Auction entity.Auction

func (s AuctionService) GetAll() ([]Auction, error) {
	db := db.GetDB()
	var auction []Auction

	if err := db.Find(&auction).Error; err != nil {
		return nil, err
	}

	return auction, nil
}

func (s AuctionService) CreateModel(c *gin.Context) (Auction, error) {
	db := db.GetDB()
	var auction Auction

	if err := c.BindJSON(&auction); err != nil {
		return auction, err
	}

	if err := db.Create(&auction).Error; err != nil {
		return auction, err
	}

	return auction, nil
}

func (s AuctionService) GetByID(id string) (Auction, error) {
	db := db.GetDB()
	var auction Auction

	if err := db.Where("id = ?", id).First(&auction).Error; err != nil {
		return auction, err
	}

	return auction, nil
}

func (s AuctionService) UpdateByID(id string, c *gin.Context) (Auction, error) {
	db := db.GetDB()
	var auction Auction

	if err := db.Where("id = ?", id).First(&auction).Error; err != nil {
		return auction, err
	}

	if err := c.BindJSON(&auction); err != nil {
		return auction, err
	}

	db.Save(&auction)

	return auction, nil
}

func (s AuctionService) DeleteByID(id string) error {
	db := db.GetDB()
	var auction Auction

	if err := db.Where("id = ?", id).Delete(&auction).Error; err != nil {
		return err
	}

	return nil
}
