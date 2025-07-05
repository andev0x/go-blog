package model

type Comment struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	PostID  uint   `json:"post_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Rating  int    `json:"rating"`
}
