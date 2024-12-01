package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"tutorme.com/api"
	"tutorme.com/repository"
)

func run() error {
	// Loading in env variables
	url := os.Getenv("TURSO_DATABASE_URL")
	token := os.Getenv("TURSO_AUTH_TOKEN")
	port := ":" + os.Getenv("DOCKER_PORT")
	flag.Parse()

	// Config logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//Concatenated database url
	connUri := fmt.Sprintf("%s?authToken=%s", url, token)

	db, err := sql.Open("libsql", connUri)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repository.New(db)

	slog.Info("Starting Server...", "Port", port)

	fiberCfg := fiber.Config{
		ReadTimeout:           time.Second * 5,
		WriteTimeout:          time.Second * 10,
		IdleTimeout:           time.Second * 60,
		DisableStartupMessage: true,
	}

	server := api.New(fiberCfg, port, logger, queries)

	return server.Start()
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
