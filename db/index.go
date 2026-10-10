package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn
func InitDB() {
	url := "postgres://postgres:adminPassword@localhost:5434/todo-tasks?sslmode=disable"
	var err error
	DB, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Error in connecting DB!")
		os.Exit(1)
	}

	log.Printf("Connected with DB")
	
}