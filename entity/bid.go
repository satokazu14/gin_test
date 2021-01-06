package entity

type Bid struct {
	AuctionId int    `json:"auction_id"`
	UserId    int    `json:"user_id"`
	CarId     int    `json:"car_id"`
	Price     int    `json:"price"`
	BidTime   string `json:"bid_time"`
}
