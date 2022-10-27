package main

import (
	"log"

	"github.com/mnarvaezm96/doom_go/db"
	"github.com/mnarvaezm96/doom_go/handlers"
)

func main() {

	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")

	}

	handlers.Manejadores()

}
