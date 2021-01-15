package entity

type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	PassWord     string `json:"password"`
	Mail         string `json:"mail"`
	Phone        string `json:"phone"`
	Birthday     string `json:"birthday"`
	Zip          string `json:"zip"`
	Address      string `json:"address"`
	CardNumber   string `json:"card_number"`
	CardDate     string `json:"card_date"`
	CardSecurity string `json:"card_security"`
}
