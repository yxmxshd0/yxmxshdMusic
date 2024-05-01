package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"yxmxshdMusic/config"
)

type mongoDB struct {
	cfg *config.MongoDB
}

type MongoDB interface {
	NewMongoDBConnection(ctx context.Context) (*mongo.Database, error)
}

func New(cfg config.MongoDB) MongoDB {
	return &mongoDB{cfg: &cfg}
}

func (m *mongoDB) NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	URI := fmt.Sprintf("mongodb://%s:%s", m.cfg.Host, m.cfg.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB connected")

	return client.Database("yxmxshdMusic"), nil
}
