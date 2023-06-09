package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Molizane/gofinance-backend/api"
	db "github.com/Molizane/gofinance-backend/db/sqlc"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Router.Run(serverAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
		return
	}
}
