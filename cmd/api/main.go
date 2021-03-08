package main

import (
	"log"

	"github.com/wilgustavo/turnos/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
