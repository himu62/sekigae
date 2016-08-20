package app

func (app *App) Init() {
	app.Engine.StaticFile("/", "public/index.html")
	app.Engine.StaticFile("/favicon.ico", "public/favicon.ico")
	app.Engine.StaticFile("/css/style.css", "public/css/style.css")
	app.Engine.StaticFile("/js/bundle.js", "public/js/bundle.js")
}
