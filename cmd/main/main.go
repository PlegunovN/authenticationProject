package main

import (
	"authenticationProject/internal/logger"
	"authenticationProject/internal/server"
	"authenticationProject/internal/users"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	sLogger := logger.InitLogger()
	defer sLogger.Sync()

	db, err := sqlx.Connect("postgres", "host=localhost port=5434 user=post password=1234 dbname=authbook sslmode=disable")
	if err != nil {
		sLogger.Fatalf("not connected to db: %w", err)
	}
	storage := users.New(db, sLogger)
	server.AuthStart(storage, sLogger)

}
