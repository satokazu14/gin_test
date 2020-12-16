package entity

type Notification struct {
	ID       uint   `json:"id"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	DeleteAt string `json:"delete_at"`
}
