package handler

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"technical-task/pkg/models"
)

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.Product
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Unable to read data from request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, "Invalid input body", http.StatusBadRequest)
		return
	}

	id, err := h.services.Product.CreateProduct(input)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"id": id}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error when encoding object to JSON", http.StatusInternalServerError)
		return
	}
}

type GetProduct struct {
	ID string `json:"id,omitempty" bson:"_id,omitempty"`
}

func (h *Handler) getProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product GetProduct
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Unable to read data from request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &product); err != nil {
		http.Error(w, "Invalid input body", http.StatusBadRequest)
		return
	}

	input, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		http.Error(w, "Failed to get ObjectID from Hex", http.StatusBadRequest)
		return
	}

	resProduct, err := h.services.Product.GetProductByID(input)
	if err != nil {
		http.Error(w, "Failed to get product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"product": resProduct}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error when encoding object to JSON", http.StatusInternalServerError)
		return
	}
}

type UpdateProduct struct {
	ID      string         `json:"id,omitempty" bson:"_id,omitempty"`
	Product models.Product `json:"product"`
}

func (h *Handler) updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input UpdateProduct
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Unable to read data from request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, "Invalid input body", http.StatusBadRequest)
		return
	}

	in := models.Product{
		Name:        input.Product.Name,
		Description: input.Product.Description,
		Price:       input.Product.Price,
		SellerID:    input.Product.SellerID,
	}

	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		http.Error(w, "Failed to get ObjectID from Hex", http.StatusBadRequest)
		return
	}

	err = h.services.Product.UpdateProduct(id, in)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"update": http.StatusAccepted}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error when encoding object to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product GetProduct
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Unable to read data from request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &product); err != nil {
		http.Error(w, "Invalid input body", http.StatusBadRequest)
		return
	}

	input, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		http.Error(w, "Failed to get ObjectID from Hex", http.StatusBadRequest)
		return
	}

	count, err := h.services.Product.DeleteProduct(input)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"count": count}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error when encoding object to JSON", http.StatusInternalServerError)
		return
	}
}
