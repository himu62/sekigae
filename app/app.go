package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type (
	App struct {
		Engine *gin.Engine
	}
)

func New() *App {
	return &App{Engine: gin.Default()}
}

func (app *App) Run() {
	app.Engine.Run(":8888")
}
