package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	userCollections    = "user"
	productCollections = "product"
	//sellerCollections   = "seller"
	//customerCollections = "customer"
	//orderCollections    = "order"
)

var ctx = context.Background()

func NewClient(username, password, host, port, database string) (*mongo.Database, error) {
	var atlasURI string

	if username == "" || password == "" {
		atlasURI = fmt.Sprintf("mongodb://%s:%s/", host, port)
	} else {
		atlasURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/", username, password, host, port)
	}

	clientOptions := options.Client().ApplyURI(atlasURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(database), nil
}
