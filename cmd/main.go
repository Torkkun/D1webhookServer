package main

import (
	"app/pkg/server"
)

func main() {
	router := server.NewRouter()
	router.Run()
}
