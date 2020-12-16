package entity

type Bid struct {
	ID        uint   `json:"id"`
	AuctionId int    `json:"auction_id"`
	UserId    int    `json:"user_id"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
	DeleteAt  string `json:"delete_at"`
}
