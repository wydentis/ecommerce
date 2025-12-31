package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/wydentis/ecommerce/internal/env"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	log.Printf("connected to db, dsn: %s", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.Run(api.Mount()); err != nil {
		log.Printf("server has failed to start. err: %s", err)
		os.Exit(1)
	}
}
