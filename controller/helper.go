package controller

import (
	"github.com/gin-gonic/gin"
)

func error(c *gin.Context, status int, msg string) {
	c.JSON(status, map[string]string{"error": msg})
}
