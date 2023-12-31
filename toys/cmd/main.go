package main

import (
	"log"
	"net/http"

	"github.com/jwkidd3/microservices_final/toys/internal/routes"
)

func main() {
	r := routes.Handlers()

	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatal(err)
	}
}
