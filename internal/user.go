package internal

import (
	"burger.local/models"
)

func InsertUser(user models.User) error {
	_, err := Conn.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password) // Hash the password
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserUsername(user models.User) error {
	_, err := Conn.Exec("UPDATE users SET username = ? WHERE id = ?", user.Username, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserPassword(user models.User) error {
	_, err := Conn.Exec("UPDATE users SET password = ? WHERE id = ?", user.Password, user.ID) // Hash the password
	if err != nil {
		return err
	}

	return nil
}

// Used to delete a user
func DeleteUser(user models.User) error {
	_, err := Conn.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		return err
	}

	return nil
}

// Used to get a user's information
func GetUserByID(id int) (models.User, error) {
	var user models.User
	err := Conn.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

// This will only be used to check if a user exists or get the user's ID
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := Conn.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
