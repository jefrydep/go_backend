package main

import (
	"log"

	"github.com/jefrydep/go_backend/bd"
	"github.com/jefrydep/go_backend/handlers"
)

func main() {
	if bd.CheckedConnection() == 0 {
		log.Fatal("sin conexion ala bd")
		return
	}
	handlers.Manejadores()

}
