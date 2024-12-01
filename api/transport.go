package api

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"tutorme.com/repository"
)

type Storer interface {
	CreateDatabaseDetail(ctx context.Context, arg repository.CreateDatabaseDetailParams) error
	CreateSchool(ctx context.Context, arg repository.CreateSchoolParams) error
}

type Server struct {
	server *fiber.App
	port   string
	logger *slog.Logger
	db     Storer
}

func New(config fiber.Config, port string, logger *slog.Logger, db Storer) *Server {
	engine := &Server{
		server: fiber.New(config),
		port:   port,
		logger: logger,
		db:     db,
	}

	// Register Middlewares
	engine.server.Use(fiberlogger.New())

	// Register Routes
	engine.register()

	return engine
}

func (s *Server) Start() error {
	return s.server.ListenTLS(s.port, "certs/cert.pem", "certs/key.pem")
}
