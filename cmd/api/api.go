package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/valievdaler/task-tracker/internal/store"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router){
		r.Get("/health",app.healthCheckHandler)
	})

	
	
	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
