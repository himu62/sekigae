package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/himu62/sekigae/db"
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
		g.GET("/list", ctl.list)
		g.POST("", ctl.create)
		g.POST("/image", ctl.uploadImage)
		g.PUT("/:id", ctl.update)
		g.DELETE("/:id", ctl.delete)
	}
}

func (ctl *User) list(c *gin.Context) {
	db, err := db.New()
	if err != nil {
		error(c, 500, "failed to connect DB.")
		return
	}
	data, err := model.FindUsers(db)
	if err != nil {
		error(c, 500, "failed to fetch users.")
		return
	}

	c.JSON(200, data)
}

func (ctl *User) create(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		error(c, 400, "request parameter is wrong.")
		return
	}
	db, err := db.New()
	if err != nil {
		error(c, 500, "failed to connect DB.")
		return
	}
	if err := user.Insert(db); err != nil {
		error(c, 500, "failed to insert record.")
		return
	}

	c.JSON(200, user)
}

func (ctl *User) uploadImage(c *gin.Context) {
	_, header, err := c.Request.FormFile("image")
	if err != nil {
		error(c, 400, "request file is invalid.")
		return
	}

	url, err := model.CreateImage(header)
	if err != nil {
		error(c, 500, "failed to upload file.")
		return
	}

	data := map[string]string{"file": url}
	c.JSON(201, data)
}

func (ctl *User) update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		error(c, 400, "request parameter is wrong.")
		return
	}
	user := &model.User{}
	if err = c.Bind(user); err != nil {
		error(c, 400, "request parameter is wrong.")
		return
	}
	user.ID = uint8(id)
	db, err := db.New()
	if err != nil {
		error(c, 500, "failed to connect DB.")
		return
	}
	if err := user.Update(db); err != nil {
		error(c, 500, "failed to update record.")
		return
	}

	c.JSON(200, user)
}

func (ctl *User) delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		error(c, 400, "request parameter is wrong.")
		return
	}
	user := &model.User{ID: uint8(id)}
	db, err := db.New()
	if err != nil {
		error(c, 500, "failed to connect DB.")
		return
	}
	if err := user.Delete(db); err != nil {
		error(c, 500, "failed to delete record.")
		return
	}

	data := map[string]interface{}{"id": 0}
	c.JSON(200, data)
}
