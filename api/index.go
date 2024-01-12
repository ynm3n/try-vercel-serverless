package handler

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dsn := os.Getenv("POSTGRES_URL_NON_POOLING")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}

	fmt.Fprintln(io.MultiWriter(w, os.Stdout), "connection and ping is ok", time.Now())
}
