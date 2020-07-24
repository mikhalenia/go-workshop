package main

import (
	"server"
)

func start() {
	server.InitServer()
	server.StartWebServer()
}
