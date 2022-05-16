package main

import (
	"database/sql"
	"log"

	"github.com/Heartbeat1608/simplebank/api"
	db "github.com/Heartbeat1608/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:5000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Default().Printf("[ERR] %v", err)
		log.Fatal("cannot connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalf("Cannot start server %v", err)
	}
}
