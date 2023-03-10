package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"technical-task/pkg/models"
	"technical-task/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) CreateProduct(product models.Product) (string, error) {
	return p.repo.CreateProduct(product)
}

func (p *ProductService) GetProductByID(id primitive.ObjectID) (product models.Product, err error) {
	return p.repo.GetProductByID(id)
}

func (p *ProductService) UpdateProduct(id primitive.ObjectID, data models.Product) (err error) {
	return p.repo.UpdateProduct(id, data)
}

func (p *ProductService) DeleteProduct(id primitive.ObjectID) (count int64, err error) {
	return p.repo.DeleteProduct(id)
}
