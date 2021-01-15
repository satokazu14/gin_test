package entity

type Favorite struct {
	ID      uint `json:"id"`
	UserId  int  `json:"user_id"`
	GenreId int  `json:"genre_id"`
}
