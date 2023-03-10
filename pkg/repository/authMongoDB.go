package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"technical-task/pkg/models"
)

type AuthMongoDB struct {
	db *mongo.Database
}

func NewAuthMongoDB(db *mongo.Database) *AuthMongoDB {
	return &AuthMongoDB{db: db}
}

func (r *AuthMongoDB) CreateUser(user models.User) (string, error) {
	collections := r.db.Collection(userCollections)

	result, err := collections.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	id := insertedID.Hex()

	return id, nil
}

func (r *AuthMongoDB) GetUser(username, password string) (models.User, error) {
	var user models.User
	collections := r.db.Collection(userCollections)
	filter := bson.M{
		"username": username,
		"password": password,
	}

	result := collections.FindOne(ctx, filter)
	if result.Err() != nil {
		return models.User{}, fmt.Errorf("failed to find user by ID")
	}
	if err := result.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}
