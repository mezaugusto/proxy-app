package server

import (
	"os"

	"github.com/kataras/iris"
)

// SetUp returns the application instance
func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	return app
}

// RunServer runs the application
func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(os.Getenv("PORT")),
	)
}
