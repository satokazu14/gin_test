package entity

type Auction struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StartDay string `json:"start_day"`
	EndFlg   int    `json:"end_flg"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	DeleteAt string `json:"delete_at"`
	Cars     []Car  `gorm:"many2many:auction_cars;"`
}
