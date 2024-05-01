package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	
	"yxmxshdMusic/internal/models"
)

type MongoRepos interface {
	SendDocument(ctx *context.Context, send models.DocumentsToSend) error
}

type repo struct {
	db *mongo.Database
}

func New(db *mongo.Database) MongoRepos {
	return &repo{db: db}
}

func (r *repo) SendDocument(ctx *context.Context, send models.DocumentsToSend) error {
	collection := r.db.Collection("music")
	_, err := collection.InsertOne(*ctx, send)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
