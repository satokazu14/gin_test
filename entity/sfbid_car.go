package entity

type SfbidCar struct {
	CarID    uint   `json:"car_id"`
	UserId   int    `json:"user_id"`
	Price    int    `json:"peice"`
	Status   int    `json:"status"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	DeleteAt string `json:"delete_at"`
}
