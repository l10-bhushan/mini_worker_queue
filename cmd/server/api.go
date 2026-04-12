package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type config struct {
	addr string
}

type application struct {
	cfg *config
}

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server up and running at PORT: 8080")
	})

	return router
}

func (app *application) run(router http.Handler) error {
	server := &http.Server{
		Addr:         app.cfg.addr,
		Handler:      router,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
	}

	return server.ListenAndServe()
}
