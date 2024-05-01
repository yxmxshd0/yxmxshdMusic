package main

import (
	"context"
	"log"

	"yxmxshdMusic/config"
	"yxmxshdMusic/internal/app"
)

func main() {

	ctx := context.Background()
	cfg := config.LoadConfig()

	application := app.New(cfg)
	err := application.Run(ctx)
	if err != nil {
		log.Panic(err)
	}

}
