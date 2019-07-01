package main

import (
	"github.com/mezaugusto/proxy-app/api/handlers"
	"github.com/mezaugusto/proxy-app/api/server"
	"github.com/mezaugusto/proxy-app/api/utils"
)

func main() {
	/*
		Router Iris
		Env vars
	*/

	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
