package main

import (
	"log"

	"github.com/valievdaler/task-tracker/internal/env"
	"github.com/valievdaler/task-tracker/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Printf("Server has started at %s", app.config.addr)
	mux := app.mount()
	log.Fatal(app.run(mux))
}
