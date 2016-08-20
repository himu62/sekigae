package model

import (
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
