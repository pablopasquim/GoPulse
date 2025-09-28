package main

import (
	"fmt"
	"log"

	"github.com/pablopasquim/GoPulse/routes"
)

func main() {
	fmt.Println("Starting server...")
	log.Println("Servidor rodando em http://localhost:8080")
	routes.HandleRequest()
}
