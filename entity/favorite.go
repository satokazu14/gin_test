package entity

type Favorite struct {
	ID       uint   `json:"id"`
	UserId   int    `json:"user_id"`
	GenreId  int    `json:"genre_id"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	DeleteAt string `json:"delete_at"`
}
