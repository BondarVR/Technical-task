package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"technical-task/pkg/models"
)

type ProductMongoDB struct {
	db *mongo.Database
}

func NewProductMongoDB(db *mongo.Database) *ProductMongoDB {
	return &ProductMongoDB{db: db}
}

func (p *ProductMongoDB) CreateProduct(product models.Product) (string, error) {
	collections := p.db.Collection(productCollections)

	result, err := collections.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	id := insertedID.Hex()

	return id, nil
}

func (p *ProductMongoDB) GetProductByID(id primitive.ObjectID) (product models.Product, err error) {
	collections := p.db.Collection(productCollections)

	filter := bson.M{"_id": id}
	result := collections.FindOne(ctx, filter)
	if result.Err() != nil {
		return product, fmt.Errorf("failed to find product by ID")
	}

	if err := result.Decode(&product); err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductMongoDB) UpdateProduct(id primitive.ObjectID, data models.Product) (err error) {
	collections := p.db.Collection(productCollections)
	filter := bson.M{"_id": id}

	userBytes, err := bson.Marshal(data)
	if err != nil {
		return err
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := collections.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("failed to update product")
	}

	return nil
}

func (p *ProductMongoDB) DeleteProduct(id primitive.ObjectID) (count int64, err error) {
	collections := p.db.Collection(productCollections)
	filter := bson.M{"_id": id}

	result, err := collections.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	if result.DeletedCount == 0 {
		return 0, fmt.Errorf("failed to delete and product by ID")
	}
	return result.DeletedCount, nil
}
