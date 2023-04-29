package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Molizane/gofinance-backend/api"
	db "github.com/Molizane/gofinance-backend/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
		return
	}

	fmt.Println("Server running at: ", serverAddress)
}
