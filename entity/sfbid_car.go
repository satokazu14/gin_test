package entity

type SfbidCar struct {
	CarID  int     `json:"car_id"`
	UserId int     `json:"user_id"`
	Price  float64 `json:"peice"`
	Status int     `json:"status"`
}
