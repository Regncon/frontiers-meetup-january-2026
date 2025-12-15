package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	root "github.com/Regncon/frontiers-meetup-january-2026/pages/root"
	"github.com/delaneyj/toolbelt"
	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	natsserver "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go/jetstream"
	_ "modernc.org/sqlite"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, dbErr := sql.Open("sqlite", "presentation.db")
	if dbErr != nil {
		logger.Error("failed to open DB", slog.String("error", dbErr.Error()))
		return
	}
	_, pragmaErr := db.Exec(`
		PRAGMA journal_mode = WAL;
		PRAGMA synchronous = NORMAL;
		PRAGMA busy_timeout = 5000;
		`)
	if pragmaErr != nil {
		logger.Error("failed to set pragmas", slog.String("error", pragmaErr.Error()))
		return
	}

	if pingErr := db.Ping(); pingErr != nil {
		logger.Error("failed to ping DB", slog.String("error", pingErr.Error()))
		return
	}
	defer db.Close()

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	sessionStore := sessions.NewCookieStore([]byte("session-secret"))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(24 * time.Hour / time.Second),
		HttpOnly: true,
	}

	natsPort, err := toolbelt.FreePort()
	if err != nil {
		logger.Error("failed to get free port for nats", slog.String("error", err.Error()))
		return
	}

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	rootCtx := context.Background()

	ns, err := embeddednats.New(rootCtx, embeddednats.WithNATSServerOptions(&natsserver.Options{
		JetStream: true,
		Port:      natsPort,
	}))
	if err != nil {
		panic(fmt.Sprintf("failed to start embedded nats server: %v", err))
	}

	nc, err := ns.Client()
	if err != nil {
		panic(fmt.Sprintf("failed to create nats client: %v", err))
	}

	js, err := jetstream.New(nc)
	if err != nil {
		panic(fmt.Sprintf("failed to create jetstream context: %v", err))
	}

	kv, err := js.CreateOrUpdateKeyValue(rootCtx, jetstream.KeyValueConfig{
		Bucket:      "presentation",
		Description: "Frontiers Meetup Presentation Bucket",
		Compression: true,
		TTL:         time.Hour,
		MaxBytes:    16 * 1024 * 1024,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create or update key value store: %v", err))
	}

	root.RootLayoutRoute(router, db, sessionStore, kv, logger)

	address := ":8080"

	logger.Info("starting HTTP server", slog.String("address", address))

	httpServerError := http.ListenAndServe(address, router)
	if httpServerError != nil {
		logger.Error("HTTP server error", slog.String("error", httpServerError.Error()))
	}
}
