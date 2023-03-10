package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type Seller struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Phone string `json:"phone,omitempty" bson:"phone,omitempty"`
}

type Product struct {
	ID          string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string  `json:"name,omitempty" bson:"name,omitempty"`
	Description string  `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64 `json:"price,omitempty" bson:"price,omitempty"`
	SellerID    string  `json:"sellerId,omitempty" bson:"seller_id,omitempty"`
}

type Customer struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Phone string `json:"phone,omitempty" bson:"phone,omitempty"`
}

type Order struct {
	ID         string   `json:"id,omitempty" bson:"_id,omitempty"`
	CustomerID string   `json:"customerId,omitempty" bson:"customer_id,omitempty"`
	Products   []string `json:"products,omitempty" bson:"products,omitempty"`
}
