package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"yxmxshdMusic/config"
	"yxmxshdMusic/internal/databases/repos"
	"yxmxshdMusic/internal/handlers"
)

type Server struct {
	cfg   *config.Server
	mongo repos.MongoRepos
}

func New(cfg config.Server, r repos.MongoRepos) *Server {
	return &Server{
		cfg:   &cfg,
		mongo: r,
	}
}

func (s *Server) Run(ctx *context.Context) error {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	routes := handlers.New(s.mongo, ctx)
	routes.RegisterRoutes(api)

	serverAddress := fmt.Sprintf("%s:%v", s.cfg.Host, s.cfg.Port)

	return app.Listen(serverAddress)

}
