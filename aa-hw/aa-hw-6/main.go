package main

import (
	tcp "mytcp"
	"os"
)

func main() {
	args := os.Args
	port := ""

	if len(args) > 1 {
		port = args[1]
	} else {
		port = "8080"
	}

	tcp.Listen(port)
}
