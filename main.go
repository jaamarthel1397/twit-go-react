package main

import (
	"log"

	"github.com/jaamarthel1397/twit-go-react/bd"
	"github.com/jaamarthel1397/twit-go-react/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
