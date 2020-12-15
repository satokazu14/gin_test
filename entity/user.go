package entity

type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Mail         string `json:"mail"`
	Phone        int    `json:"phone"`
	Birthday     string `json:"birthday"`
	Zip          int    `json:"zip"`
	Address      string `json:"address"`
	CardNumber   int    `json:"card_number"`
	CardDate     string `json:"card_date"`
	CardSecurity int    `json:"card_security"`
	CreateAt     string `json:"create_at"`
	UpdateAt     string `json:"update_at"`
	DeleteAt     string `json:"delete_at"`
}
