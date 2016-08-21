package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/himu62/sekigae/model"
)

type (
	// User controller
	User struct{}
)

// AddUserRoutes make routes of User controller
func AddUserRoutes(r *gin.RouterGroup) {
	ctl := &User{}

	g := r.Group("/users")
	{
		g.GET("/", ctl.fetch)
		g.GET("/list", ctl.list)
		g.POST("/", ctl.create)
	}
}

func (ctl *User) fetch(c *gin.Context) {
	data := &model.User{}
	c.JSON(200, data)
}

func (ctl *User) list(c *gin.Context) {
	data := make([]model.User, 5)
	c.JSON(200, data)
}

func (ctl *User) create(c *gin.Context) {
	data := &model.User{}
	c.JSON(200, data)
}
