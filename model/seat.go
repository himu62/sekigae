package model

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/himu62/sekigae/model/request"
)

type (
	// Seats is seats table in the class room
	Seat struct {
		ID      uint32
		Created time.Time
		Data    []User
	}
)

// NewSeats creates Seats instance initialized by request
func NewSeat(c *gin.Context) (*Seat, error) {
	// クライアントからのリクエストを、リクエストモデルにバインドする
	req := &request.Seat{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	seat := &Seat{
		Created: now(),
		Data:    make([]User, 0, 30),
	}

	// リクエストモデルをDBのモデルに変換する
	for _, v := range req.Data {
		u := User{ID: v}
		seat.Data = append(seat.Data, u)
	}

	return seat, nil
}

func FindSeats(db *sql.DB) ([]Seat, error) {
	seats := make([]Seat, 0, 30)

	rows, err := db.Query("SELECT ID, Created FROM seats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &Seat{}
		if err := rows.Scan(s.ID, s.Created); err != nil {
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

		seats = append(seats, *s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

func FindSeatByID(db *sql.DB, id uint32) (*Seat, error) {
	seat := &Seat{}
	err := db.QueryRow("SELECT ID, Created FROM seats WHERE ID=?", id).
		Scan(seat.ID, seat.Created)
	if err != nil {
		return nil, err
	}

	seat.Data = make([]User, 0, 30)

	data, err := db.Query(`
		SELECT Name, Github, Image
		FROM seats_users s
		JOIN users u ON s.UserID=u.ID
		WHERE s.SeatID=?
		`, seat.ID)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var u User
		if err := data.Scan(&u.Name, &u.Github, &u.Image); err != nil {
			return nil, err
		}
		seat.Data = append(seat.Data, u)
	}
	if err := data.Err(); err != nil {
		return nil, err
	}

	return &seat, nil
}

func (seat *Seat) Insert(db *sql.DB) error {
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

	for _, v := range seat.Data {
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

func (seat *Seat) Update(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		"DELETE FROM seats_users WHERE SeatID=?",
		seat.ID,
	); err != nil {
		tx.Rollback()
		return err
	}

	for _, v := range seat.Data {
		if _, err := tx.Exec(
			"INSERT INTO seats_users(SeatID, UserID) VALUES(?,?)",
			seat.ID,
			v.ID,
		); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (seat *Seat) Delete(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Query(
		"DELETE FROM seats WHERE ID=?",
		seat.ID,
	); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
