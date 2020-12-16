package entity

type User struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	PassWord     string `json:"passwoed"`
	Mail         string `json:"Mail"`
	Phone        string `json:"phone"`
	Birthday     string `json:"birthday"`
	Zip          string `json:"zip"`
	Address      string `json:"address"`
	CardNumber   string `json:"card_number"`
	CardDate     string `json:"card_date"`
	CardSecurity string `json:"card_security"`
	CreateAt     string `json:"create_at"`
	UpdateAt     string `json:"update_at"`
	DeleteAt     string `json:"delete_at"`
}
