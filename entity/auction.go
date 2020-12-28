package entity

type Auction struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
	DeleteAt    string `json:"delete_at"`
	CarIDs      []int
	AuctionCars []AuctionCar
}
