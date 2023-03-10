package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"technical-task/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GetUser(username, password string) (models.User, error)
}
type Product interface {
	CreateProduct(product models.Product) (string, error)
	GetProductByID(id primitive.ObjectID) (product models.Product, err error)
	UpdateProduct(id primitive.ObjectID, data models.Product) (err error)
	DeleteProduct(id primitive.ObjectID) (count int64, err error)
}
type Repository struct {
	Authorization
	Product
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongoDB(db),
		Product:       NewProductMongoDB(db),
	}
}
