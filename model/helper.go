package model

import (
	"time"
)

func now() time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Now().Add(9 * time.Hour)
	}
	return time.Now().In(jst)
}
