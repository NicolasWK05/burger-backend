package internal

import (
	"burger.local/models"
)

// Images are handled in internal/image.go

func InsertComment(comment models.Comment) error {
	_, err := Conn.Exec("INSERT INTO comments (user_id, burger_id, content) VALUES (?, ?, ?)", comment.UserID, comment.BurgerID, comment.Content)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCommentContent(comment models.Comment) error {
	_, err := Conn.Exec("UPDATE comments SET content = ? WHERE id = ?", comment.Content, comment.ID)
	if err != nil {
		return err
	}

	return nil
}

// Used to delete a comment
func DeleteComment(comment models.Comment) error {
	_, err := Conn.Exec("DELETE FROM comments WHERE id = ?", comment.ID)
	if err != nil {
		return err
	}

	return nil
}

// Used to get a comment's information
func GetCommentByID(id int) (models.Comment, error) {
	var comment models.Comment
	err := Conn.QueryRow("SELECT * FROM comments WHERE id = ?", id).Scan(&comment.ID, &comment.UserID, &comment.BurgerID, &comment.Content)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func GetCommentsByUserID(id int) ([]models.Comment, error) {
	var comments []models.Comment
	rows, err := Conn.Query("SELECT * FROM comments WHERE user_id = ?", id)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.BurgerID, &comment.Content)
		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func GetCommentsByBurgerID(id int) ([]models.Comment, error) {
	var comments []models.Comment
	rows, err := Conn.Query("SELECT * FROM comments WHERE burger_id = ?", id)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.BurgerID, &comment.Content)
		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func GetCommentsByUserIDAndBurgerID(userID, burgerID int) ([]models.Comment, error) {
	var comments []models.Comment
	rows, err := Conn.Query("SELECT * FROM comments WHERE user_id = ? AND burger_id = ?", userID, burgerID)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.BurgerID, &comment.Content)
		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
