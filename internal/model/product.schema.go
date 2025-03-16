package model

import (
	"time"
)

// collection of products
const (
	DatabaseProduct      = "go_ecommerce"
	ProductCollection    = "products"
	ElectronicCollection = "electronics"
	ClothingCollection   = "clothings"
)

// enum Product_Type
const (
	Product_Type_Electronic = "electronic"
	Product_Type_Clothing   = "clothing"
	Product_Type_Furniture  = "furniture"
)

// product schema mongo
type ProductSchema struct {
	Product_Name        string      `bson:"product_name" validate:"required"`
	Product_Thumb       string      `bson:"product_thumb" validate:"required"`
	Product_Description string      `bson:"product_description,omitempty"`
	Product_Price       float64     `bson:"product_price" validate:"required"`
	Product_Quantity    int         `bson:"product_quantity" validate:"required"`
	Product_Type        string      `bson:"product_type" validate:"required"` // type is electronic, clothing, furniture, ...
	Product_Shop        string      `bson:"product_shop" validate:"required"`
	Product_Atributes   interface{} `bson:"product_attributes" validate:"required"`
	Product_CreatedAt   time.Time   `bson:"product_created_at" validate:"required"`
	Product_UpdatedAt   time.Time   `bson:"product_updated_at" validate:"required"`
	Product_DeletedAt   time.Time   `bson:"product_deleted_at,omitempty"`
}

// clothing schema definition
type ClothingSchema struct {
	Brand     string    `bson:"brand" validate:"required"`
	Size      string    `bson:"size" validate:"required"`
	Material  string    `bson:"metarial" validate:"required"`
	CreatedAt time.Time `bson:"product_created_at" validate:"required"`
	UpdatedAt time.Time `bson:"product_updated_at" validate:"required"`
	DeletedAt time.Time `bson:"product_deleted_at,omitempty"`
}

// Electronic schema definition
type ElectronicSchema struct {
	Manufacturer string    `bson:"manufacturer" validate:"required"`
	Model        string    `bson:"model" validate:"required"`
	Color        string    `bson:"color" validate:"required"`
	CreatedAt    time.Time `bson:"product_created_at" validate:"required"`
	UpdatedAt    time.Time `bson:"product_updated_at" validate:"required"`
	DeletedAt    time.Time `bson:"product_deleted_at,omitempty"`
}
