package handlers

import (
	"encoding/json"

	"github.com/kataras/iris"
	"github.com/mezaugusto/proxy-app/api/middleware"
)

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application) {
	app.Get("/", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context) {
	response, err := json.Marshal(middleware.Q)
	if err != nil {
		c.JSON(iris.Map{"status": 400, "result": "parse error"})
		return
	}
	c.JSON(iris.Map{"result": string(response)})
}
