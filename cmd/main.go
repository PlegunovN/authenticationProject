package main

import (
	"fmt"
	"github.com/PlegunovN/authenticationProject/configs"
	"github.com/PlegunovN/authenticationProject/internal/logger"
	"github.com/PlegunovN/authenticationProject/internal/server"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	sLogger := logger.InitLogger()
	defer sLogger.Sync()

	cfg, err := configs.LoadConfig("./.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode))
	if err != nil {
		sLogger.Fatalf("not connected to db: %w", err)
	}

	storage := users.New(db, sLogger)
	server.AuthStart(storage, sLogger)

}
