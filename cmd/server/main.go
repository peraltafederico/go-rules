package main

import (
	"log"
	"net/http"

	"github.com/peraltafederico/go-rules/internal/routes"
)

func main() {
	r := routes.InitRouter()

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
