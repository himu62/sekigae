package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/himu62/sekigae/db"
	"github.com/himu62/sekigae/model"
)

type (
	// Seat controller
	Seat struct{}
)

// AddSeatRoutes make routes of seat controller
func AddSeatRoutes(r *gin.RouterGroup) {
	ctl := &Seat{}

	g := r.Group("/seats")
	{
		g.GET("/:id", ctl.fetch)
		g.GET("/list", ctl.list)
		g.POST("", ctl.create)
	}
}

func (ctl *Seat) fetch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(400, NewError("request parameter is wrong."))
		return
	}
	db, err := db.New()
	if err != nil {
		c.JSON(500, NewError("failed to connect DB"))
		return
	}
	data, err := model.FindSeatsByID(db, uint32(id))
	if err != nil {
		c.JSON(500, NewError("failed to fetch seats"))
		return
	}

	c.JSON(200, data)
}

func (ctl *Seat) list(c *gin.Context) {
	db, err := db.New()
	if err != nil {
		c.JSON(500, NewError("failed to connect DB"))
		return
	}
	data, err := model.FindSeats(db)
	if err != nil {
		c.JSON(500, NewError("failed to fetch seats"))
		return
	}

	c.JSON(200, data)
}

func (ctl *Seat) create(c *gin.Context) {
	data := &model.Seats{}
	c.JSON(200, data)
}
