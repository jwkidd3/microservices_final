package main

import (
	"log"
	"net/http"

	"github.com/jwkidd3/microservices_final/inventory/internal/routes"
)

func main() {
	r := routes.Handlers()

	err := http.ListenAndServe(":5100", r)
	if err != nil {
		log.Fatal(err)
	}
}
