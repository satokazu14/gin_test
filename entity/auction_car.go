package entity

type AuctionCar struct {
	AuctionId int    `json:"auction_id"`
	CarId     int    `json:"car_id"`
	StartTime string `json:"start_time"`
}
