package models

type Comment struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	BurgerID int    `json:"burger_id" db:"burger_id"`
	Content  string `json:"content" db:"content"`
}
