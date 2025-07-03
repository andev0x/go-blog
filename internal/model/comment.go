package model

type Comment struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	PostID  uint   `json:"post_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}
