package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dsn := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("pgx", dsn)
	log.Println(err)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	err = db.PingContext(ctx)
	log.Println(err)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "connection and ping is ok")
}
