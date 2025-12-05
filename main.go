package main

//module github.com/Regncon/frontiers-meetup-january-2026
import (
	"log"
	"net/http"

	root "github.com/Regncon/frontiers-meetup-january-2026/pages/root"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", rootIndexHandler)

	address := ":8080"

	log.Printf("Server listening on %s", address)

	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}
}

func rootIndexHandler(w http.ResponseWriter, r *http.Request) {
	component := root.RootIndex()

	handler := templ.Handler(component)

	handler.ServeHTTP(w, r)
}
