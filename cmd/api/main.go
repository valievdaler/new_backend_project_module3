package main

import (
	"log"

	"github.com/valievdaler/task-tracker/internal/db"
	"github.com/valievdaler/task-tracker/internal/env"
	"github.com/valievdaler/task-tracker/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5433/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_OPEN_TIME", "15m"),
		},
	}
	db,err:=db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err!=nil{
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Printf("Server has started at %s", app.config.addr)
	mux := app.mount()
	log.Fatal(app.run(mux))
}
