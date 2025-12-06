package main

import (
	"log"
	"net/http"
	"time"

	root "github.com/Regncon/frontiers-meetup-january-2026/pages/root"
	"github.com/delaneyj/toolbelt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	sessionStore := sessions.NewCookieStore([]byte("session-secret"))
	sessionStore.MaxAge(int(24 * time.Hour / time.Second))

	natsPort, err := toolbelt.FreePort()
	if err != nil {
		log.Fatalf("error finding free port: %v", err)
		return
	}

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	//func RootLayoutRoute(router chi.Router, natsPort int, store sessions.Store) {
	root.RootLayoutRoute(router, natsPort, sessionStore)

	address := ":8080"

	log.Printf("Server listening on %s", address)

	httpServerError := http.ListenAndServe(address, router)
	if httpServerError != nil {
		log.Fatal(err)
	}
}
