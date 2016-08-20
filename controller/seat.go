package controller

import (
	"github.com/gin-gonic/gin"
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
		g.GET("/", ctl.fetch)
		g.GET("/list", ctl.list)
		g.POST("/", ctl.create)
	}
}

func (ctl *Seat) fetch(c *gin.Context) {
	data := &model.Seats{}
	c.JSON(200, data)
}

func (ctl *Seat) list(c *gin.Context) {
	data := make([]model.Seats, 5)
	c.JSON(200, data)
}

func (ctl *Seat) create(c *gin.Context) {
	data := &model.Seats{}
	c.JSON(200, data)
}
