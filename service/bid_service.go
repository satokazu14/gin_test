package service

import (
	"gin_test/db"
	"gin_test/entity"
	"time"

	"github.com/gin-gonic/gin"
)

type BidService struct{}

type Bid entity.Bid

func (s BidService) GetAll() ([]Bid, error) {
	db := db.GetDB()
	var bid []Bid

	if err := db.Find(&bid).Error; err != nil {
		return nil, err
	}

	return bid, nil
}

func (s BidService) GetByAuctionID(auctionId string) ([]Bid, error) {
	db := db.GetDB()
	var bid []Bid

	if err := db.Where("auction_id = ?", auctionId).Find(&bid).Error; err != nil {
		return nil, err
	}

	return bid, nil
}

func (s BidService) GetByUserID(userId string) ([]Bid, error) {
	db := db.GetDB()
	var bid []Bid

	if err := db.Where("user_id = ?", userId).Find(&bid).Error; err != nil {
		return nil, err
	}

	return bid, nil
}

func (s BidService) GetByCarID(userId string) ([]Bid, error) {
	db := db.GetDB()
	var bid []Bid

	if err := db.Where("car_id = ?", userId).Find(&bid).Error; err != nil {
		return nil, err
	}

	return bid, nil
}

func (s BidService) GetResult(auctionId string) ([]Bid, error) {
	db := db.GetDB()
	var bid []Bid

	if err := db.Where("auction_id = ?", auctionId).Last(&bid).Error; err != nil {
		return nil, err
	}

	return bid, nil
}

func (s BidService) CreateBid(c *gin.Context) (Bid, error) {
	db := db.GetDB()
	var layout = "15:04:05"
	var bid Bid
	now := time.Now()
	str := now.Format(layout)

	if err := c.BindJSON(&bid); err != nil {
		return bid, err
	}

	bid.BidTime = str

	if err := db.Create(&bid).Error; err != nil {
		return bid, err
	}

	return bid, nil
}
