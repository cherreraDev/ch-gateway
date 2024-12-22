package main

import (
	"ch-gateway/cmd/api/bootstrap"
	"fmt"
	"log"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(fmt.Printf("Se ha producido un error al arrancar el servidor: %e", err))
	}
}
