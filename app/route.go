package app

import (
	"github.com/himu62/sekigae/controller"
)

func (app *App) Init() {
	app.Engine.StaticFile("/", "public/index.html")
	app.Engine.StaticFile("/favicon.ico", "public/favicon.ico")
	app.Engine.StaticFile("/css/style.css", "public/css/style.css")
	app.Engine.StaticFile("/js/bundle.js", "public/js/bundle.js")

	g := app.Engine.Group("/api", CSRF())
	controller.AddSeatRoutes(g)
	controller.AddUserRoutes(g)
}
