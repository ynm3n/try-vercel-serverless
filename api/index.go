package handler

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dsn := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
}
