package main

import (
	"log"
	"net/http"

	"github.com/microservices_final/jwkidd3/users/internal/routes"
)

func main() {
	r := routes.Handlers()

	err := http.ListenAndServe(":5050", r)
	if err != nil {
		log.Fatal(err)
	}
}
