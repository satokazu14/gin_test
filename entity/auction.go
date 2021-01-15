package entity

type Auction struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StartDay string `json:"start_day"`
	EndFlg   int    `json:"end_flg"`
	Cars     []Car  `gorm:"many2many:auction_cars;"`
}
