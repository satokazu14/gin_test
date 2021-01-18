package service

import (
	"gin_test/db"
	"gin_test/entity"

	"github.com/gin-gonic/gin"
)

type AuctionService struct{}

type Auction entity.Auction
type AuctionCar entity.AuctionCar
type AuctionRequest entity.AuctionRequest

func (s AuctionService) GetAll() ([]Auction, error) {
	db := db.GetDB()
	var auction []Auction

	if err := db.Find(&auction).Error; err != nil {
		return nil, err
	}

	return auction, nil
}

func (s AuctionService) GetTopInfo() ([]Auction, []AuctionCar, error) {
	db := db.GetDB()
	var a []Auction
	var ac []AuctionCar

	if err := db.Where("end_flg = ?", 0).First(&a).Error; err != nil {
		return nil, nil, err
	}

	if err := db.Where("auction_id = ?", a[0].ID).Find(&ac).Error; err != nil {
		return nil, nil, err
	}

	if err := db.Preload("Cars").Where("end_flg = ?", 0).First(&a).Error; err != nil {
		return nil, nil, err
	}

	return a, ac, nil
}

func (s AuctionService) GetTopAllInfo() ([]Auction, error) {
	db := db.GetDB()
	var a []Auction

	if err := db.Find(&a).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("Cars").Find(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (s AuctionService) CreateModel(c *gin.Context) (Auction, error) {
	db := db.GetDB()
	var auction Auction
	var json AuctionRequest

	if err := c.BindJSON(&json); err != nil {
		return auction, err
	}

	db.Exec("INSERT INTO auctions (name) VALUES (?)", json.Name)

	db.Last(&auction)

	for _, v := range json.AuctionCar {
		db.Exec("INSERT INTO auction_cars (auction_id,car_id,start_time) VALUES (?,?,?)", auction.ID, v.CarId, v.StartTime)
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

func (s AuctionService) GetByCarID(id string) (AuctionCar, error) {
	db := db.GetDB()
	var auctioncar AuctionCar

	if err := db.Where("car_id = ?", id).First(&auctioncar).Error; err != nil {
		return auctioncar, err
	}

	return auctioncar, nil
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
