package main

import (
	"authserver/cmd/router"
)

const PORT = 8080

func main() {
	router.RunServer()
}
