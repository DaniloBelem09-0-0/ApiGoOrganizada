package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_user_db := os.Getenv(MONGODB_USER_DB)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	fmt.Println("Conexão realizada com sucesso! cliente: ", client)

	return client.Database(mongodb_user_db), nil
}

