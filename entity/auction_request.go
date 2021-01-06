package entity

type AuctionRequest struct {
	Name       string       `json:"name"`
	AuctionCar []AuctionCar `json:"auction_car"`
}
