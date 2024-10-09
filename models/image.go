package models

type Image struct {
	ID        int    `json:"id" db:"id"`
	CommentID int    `json:"comment_id" db:"comment_id"`
	Data      []byte `json:"data" db:"data"`
	Type      string `json:"type" db:"type"`
	Name      string `json:"name" db:"name"`
}
