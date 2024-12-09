package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"tutorme.com/api"
	"tutorme.com/internal/repository"
	"tutorme.com/util"
)

func run() error {
	url, err := util.ResolveEnv("TURSO_DATABASE_URL")
	if err != nil {
		return err
	}
	token, err := util.ResolveEnv("TURSO_AUTH_TOKEN")
	if err != nil {
		return err
	}
	port, err := util.ResolveEnv("DOCKER_PORT")
	if err != nil {
		return err
	}

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

	server := api.New(fiberCfg, ":"+port, logger, queries)

	return server.Start()
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
