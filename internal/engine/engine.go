package engine

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/alwindoss/eon"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func Run(cfg *eon.Config) error {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	setupHandlers(r)
	addr := fmt.Sprintf(":%d", cfg.Port)
	err = http.ListenAndServe(addr, r)
	return err
}

func setupHandlers(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ROUTE /"))
	})
	r.Route("/eon/v1", func(v1Router chi.Router) {
		v1Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ROUTE /eon/v1"))
		})
	})
}
