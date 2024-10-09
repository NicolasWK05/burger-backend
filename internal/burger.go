package internal

import (
	"burger.local/models"
)

func InsertBurger(burger models.Burger) error {
	_, err := Conn.Exec("INSERT INTO burgers (name, address) VALUES (?, ?)", burger.Name, burger.Address)
	if err != nil {
		return err
	}

	return nil
}

func UpdateBurgerName(burger models.Burger) error {
	_, err := Conn.Exec("UPDATE burgers SET name = ? WHERE id = ?", burger.Name, burger.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateBurgerAddress(burger models.Burger) error {
	_, err := Conn.Exec("UPDATE burgers SET address = ? WHERE id = ?", burger.Address, burger.ID)
	if err != nil {
		return err
	}

	return nil
}

// Used to delete a burger
func DeleteBurger(burger models.Burger) error {
	_, err := Conn.Exec("DELETE FROM burgers WHERE id = ?", burger.ID)
	if err != nil {
		return err
	}

	return nil
}

// Used to get a burger's information
func GetBurgerByID(id int) (models.Burger, error) {
	var burger models.Burger
	err := Conn.QueryRow("SELECT * FROM burgers WHERE id = ?", id).Scan(&burger.ID, &burger.Name, &burger.Address)
	if err != nil {
		return burger, err
	}

	return burger, nil
}

func SearchBurgerByName(name string) ([]models.Burger, error) {
	var burgers []models.Burger
	rows, err := Conn.Query("SELECT * FROM burgers WHERE name like ?", name)
	if err != nil {
		return burgers, err
	}

	for rows.Next() {
		var burger models.Burger
		err = rows.Scan(&burger.ID, &burger.Name, &burger.Address)
		if err != nil {
			return burgers, err
		}

		burgers = append(burgers, burger)
	}

	return burgers, nil
}
