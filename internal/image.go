package internal

import (
	"burger.local/models"
)

func InsertImage(image models.Image) error {
	_, err := Conn.Exec("INSERT INTO images (comment_id, data, type, name) VALUES (?, ?, ?, ?)", image.CommentID, image.Data, image.Type, image.Name)
	if err != nil {
		return err
	}

	return nil
}

func UpdateImageName(image models.Image) error {
	_, err := Conn.Exec("UPDATE images SET name = ? WHERE id = ?", image.Name, image.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateImageData(image models.Image) error {
	_, err := Conn.Exec("UPDATE images SET data = ? WHERE id = ?", image.Data, image.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateImageType(image models.Image) error {
	_, err := Conn.Exec("UPDATE images SET type = ? WHERE id = ?", image.Type, image.ID)
	if err != nil {
		return err
	}

	return nil
}

// Used to delete an image
func DeleteImage(image models.Image) error {
	_, err := Conn.Exec("DELETE FROM images WHERE id = ?", image.ID)
	if err != nil {
		return err
	}

	return nil
}

// Used to get an image's information
func GetImageByID(id int) (models.Image, error) {
	var image models.Image
	err := Conn.QueryRow("SELECT * FROM images WHERE id = ?", id).Scan(&image.ID, &image.CommentID, &image.Data, &image.Type, &image.Name)
	if err != nil {
		return image, err
	}

	return image, nil
}

func GetImagesByCommentID(id int) ([]models.Image, error) {
	var images []models.Image
	rows, err := Conn.Query("SELECT * FROM images WHERE comment_id = ?", id)
	if err != nil {
		return images, err
	}

	for rows.Next() {
		var image models.Image
		err = rows.Scan(&image.ID, &image.CommentID, &image.Data, &image.Type, &image.Name)
		if err != nil {
			return images, err
		}

		images = append(images, image)
	}

	return images, nil
}
