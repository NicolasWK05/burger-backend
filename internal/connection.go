package internal

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn *sql.DB // Global database connection
)

func NewDB() error {
	var err error

	connectionString := os.Getenv("DATABASE_URL")

	Conn, err = sql.Open("mysql", connectionString) // Move this to a configuration file
	if err != nil {
		return err
	}

	return nil
}

func Init() error {
	if Conn == nil {
		println("Attempting to connect to database")
		err := NewDB()
		if err != nil {
			println("Failed to connect to database")
			return err
		}
	}

	// Create tables here
	Conn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	)`)
	Conn.Exec(`CREATE TABLE IF NOT EXISTS burgers (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		address VARCHAR(255) NOT NULL
	)`)
	Conn.Exec(`CREATE TABLE IF NOT EXISTS comments (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		burger_id INT NOT NULL,
		content TEXT NOT NULL
	)`)
	Conn.Exec(`CREATE TABLE IF NOT EXISTS images (
		id INT AUTO_INCREMENT PRIMARY KEY,
		comment_id INT NOT NULL,
		data LONGBLOB NOT NULL,
		type VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL
	)`)
	Conn.Exec(`ALTER TABLE comments ADD FOREIGN KEY (user_id) REFERENCES users(id)`)
	Conn.Exec(`ALTER TABLE comments ADD FOREIGN KEY (burger_id) REFERENCES burgers(id)`)
	Conn.Exec(`ALTER TABLE images ADD FOREIGN KEY (comment_id) REFERENCES comments(id)`)
	Conn.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS user_username ON users (username)`)

	return nil
}

// Used for clearing tables before testing. Not used in production
func TestInit() error {

	if Conn == nil {
		println("Attempting to connect to database")
		err := NewDB()
		if err != nil {
			println("Failed to connect to database")
			return err
		}
	}

	Conn.Exec("SET FOREIGN_KEY_CHECKS = 0")
	Conn.Exec("DROP TABLE users")
	Conn.Exec("DROP TABLE burgers")
	Conn.Exec("DROP TABLE comments")
	Conn.Exec("DROP TABLE images")
	Conn.Exec("SET FOREIGN_KEY_CHECKS = 1")
	err := Init()

	if err != nil {
		return err
	}

	return nil
}
