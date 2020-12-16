package entity

type AuctionCar struct {
	ID        uint `json:"id"`
	AuctionId int  `json:"auction_id"`
	CarId     int  `json:"car_id"`
}
