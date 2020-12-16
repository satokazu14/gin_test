package entity

type Genre struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	DeleteAt string `json:"delete_at"`
}
