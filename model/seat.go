package model

import (
	"database/sql"
	"time"
)

type (
	// Seats is seats table in the class room
	Seats struct {
		ID      uint32
		Created time.Time
		Data    []User
	}
)

func FindSeats(db *sql.DB) ([]Seats, error) {
	seats := make([]Seats, 0, 30)

	rows, err := db.Query("SELECT ID, Created FROM seats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Seats
		if err := rows.Scan(&s.ID, &s.Created); err != nil {
			return nil, err
		}

		s.Data = make([]User, 0, 30)

		data, err := db.Query(`
			SELECT Name, Github, Image
			FROM seats_users s
			JOIN users u ON s.UserID=u.ID
			WHERE s.SeatID=?
			`, s.ID)
		if err != nil {
			return nil, err
		}
		defer data.Close()

		for data.Next() {
			var u User
			if err := data.Scan(&u.Name, &u.Github, &u.Image); err != nil {
				return nil, err
			}
			s.Data = append(s.Data, u)
		}
		if err := data.Err(); err != nil {
			return nil, err
		}

		seats = append(seats, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

func FindSeatsByID(db *sql.DB, id uint32) (*Seats, error) {
	var seats Seats
	err := db.QueryRow("SELECT ID, Created FROM seats WHERE ID=?", id).
		Scan(&seats.ID, &seats.Created)
	if err != nil {
		return nil, err
	}

	seats.Data = make([]User, 0, 30)

	data, err := db.Query(`
		SELECT Name, Github, Image
		FROM seats_users s
		JOIN users u ON s.UserID=u.ID
		WHERE s.SeatID=?
		`, seats.ID)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var u User
		if err := data.Scan(&u.Name, &u.Github, &u.Image); err != nil {
			return nil, err
		}
		seats.Data = append(seats.Data, u)
	}
	if err := data.Err(); err != nil {
		return nil, err
	}

	return &seats, nil
}

func (seats *Seats) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("INSERT INTO seats")
	if err != nil {
		tx.Rollback()
		return err
	}

	seatID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, v := range seats.Data {
		if _, err := tx.Exec(
			"INSERT INTO seats_users(SeatID, UserID) VALUES(?,?)",
			seatID,
			v.ID,
		); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (seats *Seats) Update(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		"DELETE FROM seats_users WHERE SeatID=?",
		seats.ID,
	); err != nil {
		tx.Rollback()
		return err
	}

	for _, v := range seats.Data {
		if _, err := tx.Exec(
			"INSERT INTO seats_users(SeatID, UserID) VALUES(?,?)",
			seats.ID,
			v.ID,
		); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (seats *Seats) Delete(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Query(
		"DELETE FROM seats WHERE ID=?",
		seats.ID,
	); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
