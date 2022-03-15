package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wahidx/go-py-comm/internal"
)

func main() {
	r := chi.NewRouter()

	r.Post("/", internal.PostData)
	r.Get("/", internal.GetData)

	log.Println("Server running...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
