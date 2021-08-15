package main

import (
	"github.com/fiuskylab/buzzao-test/server"
)

func main() {
	sv := server.NewServer()

	sv.Listen(":3000")
}
