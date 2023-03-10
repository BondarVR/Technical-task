package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"technical-task/pkg/models"
	"technical-task/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Product interface {
	CreateProduct(product models.Product) (string, error)
	GetProductByID(id primitive.ObjectID) (product models.Product, err error)
	UpdateProduct(id primitive.ObjectID, data models.Product) (err error)
	DeleteProduct(id primitive.ObjectID) (count int64, err error)
}

type Service struct {
	Authorization
	Product
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Product:       NewProductService(repo.Product),
	}
}
