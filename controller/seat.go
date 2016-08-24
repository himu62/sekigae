package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/himu62/sekigae/db"
	"github.com/himu62/sekigae/model"
)

type (
	// Seats controller
	Seat struct{}
)

// AddSeatsRoutes make routes of seats controller
func AddSeatRoutes(r *gin.RouterGroup) {
	ctl := &Seat{}

	g := r.Group("/seats")
	{
		g.GET("/detail/:id", ctl.fetch)
		g.GET("/list", ctl.list)
		g.POST("", ctl.create)
	}
}

func (ctl *Seat) fetch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		error(c, 500, "equest parameter is wrong.")
		return
	}
	db, err := db.New()
	if err != nil {
		error(c, 500, "failed to fetch DB.")
		return
	}
	seat, err := model.FindSeatByID(db, uint32(id))
	if err != nil {
		error(c, 500, "failed to fetch seat.")
		return
	}

	c.JSON(200, seat)
}

func (ctl *Seat) list(c *gin.Context) {
	db, err := db.New()
	if err != nil {
		error(c, 500, "failed to connect DB.")
		return
	}
	seats, err := model.FindSeats(db)
	if err != nil {
		error(c, 500, "failed to fetch seats.")
		return
	}

	c.JSON(200, seats)
}

func (ctl *Seat) create(c *gin.Context) {
	seat, err := model.NewSeat(c)
	if err != nil {
		error(c, 400, "request parameter is wrong.")
		return
	}

	db, err := db.New()
	if err != nil {
		error(c, 500, "failed to connect DB.")
		return
	}

	if err := seat.Insert(db); err != nil {
		error(c, 500, "failed to insert record.")
		return
	}

	c.JSON(201, seat)
}
