package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {
	// TODO: pegar informações do .env
	dsn := "postgres://admin:admin@db:5432/system?sslmode=disable"

	var (
		conn *pgx.Conn
		err  error
	)

	timeout := time.After(1 * time.Minute)
	tick := time.Tick(2 * time.Second)

	for {
		select {
		case <-timeout:
			log.Fatalf("❌ Could not connect to DB after 1 minute: %v", err)
		case <-tick:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			conn, err = pgx.Connect(ctx, dsn)
			if err == nil {
				fmt.Println("✅ Connected to PostgreSQL")
				DB = conn
				return
			}

			fmt.Println("⏳ Waiting for DB to be ready...")
		}
	}
}
