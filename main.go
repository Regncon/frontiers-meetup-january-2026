package main

//module github.com/Regncon/frontiers-meetup-january-2026
import (
	"log"
	"net/http"

	root "github.com/Regncon/frontiers-meetup-january-2026/pages/root"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	root.RootLayoutRoute(router, nil)

	address := ":8080"

	log.Printf("Server listening on %s", address)

	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}
}
