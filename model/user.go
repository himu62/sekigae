package model

import (
	"database/sql"
)

type (
	// User is profile of a Treasure student
	User struct {
		ID     uint8
		Name   string
		Github string
		Image  string
	}
)

func FindUsers(db *sql.DB) ([]User, error) {
	users := make([]User, 0, 30)

	rows, err := db.Query("SELECT ID, Name, Github, Image FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Github, &u.Image); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (user *User) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Query(
		"INSERT INTO users(Name, Github, Image) VALUES(?, ?, ?)",
		user.Name,
		user.Github,
		user.Image,
	); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (user *User) Update(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Query(
		"UPDATE users SET Name = ?, Github = ?, Image = ?",
		user.Name,
		user.Github,
		user.Image,
	); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (user *User) Delete(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Query(
		"DELETE FROM users WHERE ID = ?",
		user.ID,
	); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
