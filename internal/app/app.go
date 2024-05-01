package app

import (
	"context"
	"log"

	"yxmxshdMusic/config"
	"yxmxshdMusic/internal/databases/mongodb"
	"yxmxshdMusic/internal/databases/repos"
	"yxmxshdMusic/internal/server"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (app *App) Run(ctx context.Context) error {
	m := mongodb.New(app.cfg.MongoDB)
	mongo, err := m.NewMongoDBConnection(ctx)
	if err != nil {
		log.Panic(err)
	}

	log.Println("MongoDB connected")

	repo := repos.New(mongo)

	s := server.New(app.cfg.Server, repo)

	err = s.Run(&ctx)
	if err != nil {
		log.Panic(err)
	}

	return nil
}
