package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pgx "github.com/jackc/pgx/v4"
)

func main() {
	connectionString := os.Getenv("db_uri")

	if connectionString == "" {
		connectionString = fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable",
			os.Getenv("db_user"),
			os.Getenv("db_host"),
			os.Getenv("db_port"),
			os.Getenv("db_database"),
		)
	}

	var conn *pgx.Conn
	var err error

	for {
		conn, err = pgx.Connect(context.Background(), connectionString)
		if err == nil {
			break
		}

		log.Printf("connection failed (%v); will retry...", err)
		time.Sleep(1 * time.Second)
	}
	log.Printf("connected to database")

	var value int
	if err := conn.QueryRow(context.Background(), "select 1").Scan(&value); err != nil {
		panic(err)
	}

	fmt.Printf("All done.\n")
}
