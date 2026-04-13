package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/l10-bhushan/mini_worker_queue/internal/handler"
	"github.com/l10-bhushan/mini_worker_queue/internal/repository"
	"github.com/l10-bhushan/mini_worker_queue/internal/service"
)

type Config struct {
	Addr string
}

type Application struct {
	Cfg *Config
}

func (app *Application) Mount() http.Handler {
	repository := repository.NewJobRepository()
	service := service.NewJobService(repository)
	handler := handler.NewJobHandler(service)
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server up and running at PORT: 8080")
	})
	router.Get("/jobs", handler.GetAllJobs)
	router.Post("/jobs/create", handler.CreateJob)

	return router
}

func (app *Application) Run(router http.Handler) error {
	server := &http.Server{
		Addr:         app.Cfg.Addr,
		Handler:      router,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
	}

	return server.ListenAndServe()
}
