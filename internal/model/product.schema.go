package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// collection of products
const (
	DatabaseProduct      = "go_ecommerce"
	ProductCollection    = "products"
	ElectronicCollection = "electronics"
	ClothingCollection   = "clothes"
)

// enum Product_Type
const (
	Product_Type_Electronic = "electronic"
	Product_Type_Clothing   = "clothing"
	Product_Type_Furniture  = "furniture"
)

// product schema mongo
type ProductSchema struct {
	Product_ID          primitive.ObjectID `bson:"_id"`
	Product_Name        string             `bson:"product_name" validate:"required"`
	Product_Thumb       string             `bson:"product_thumb" validate:"required"`
	Product_Description string             `bson:"product_description,omitempty"`
	Product_Price       float64            `bson:"product_price" validate:"required"`
	Product_Quantity    int                `bson:"product_quantity" validate:"required"`
	Product_Type        string             `bson:"product_type" validate:"required"` // type is electronic, clothing, furniture, ...
	Product_Shop        primitive.ObjectID `bson:"product_shop" validate:"required"`
	Product_Atributes   interface{}        `bson:"product_attributes" validate:"required"`
	Product_Created_At  time.Time          `bson:"product_created_at" validate:"required"`
	Product_Updated_At  time.Time          `bson:"product_updated_at" validate:"required"`
	Product_Deleted_At  time.Time          `bson:"product_deleted_at,omitempty"`
}

// product schema input
type ProductSchemaInput struct {
	Product_Name        string             `bson:"product_name" validate:"required"`
	Product_Thumb       string             `bson:"product_thumb" validate:"required"`
	Product_Description string             `bson:"product_description,omitempty"`
	Product_Price       float64            `bson:"product_price" validate:"required"`
	Product_Quantity    int                `bson:"product_quantity" validate:"required"`
	Product_Shop        primitive.ObjectID `bson:"product_shop" validate:"required"`
}

// validate product input
func (p *ProductSchemaInput) Validate() (bool, string) {
	// TODO: validate
	return true, ""
}

// clothing schema definition
type ClothingSchema struct {
	ID       primitive.ObjectID `bson:"_id"`
	Brand    string             `bson:"brand" validate:"required"`
	Size     string             `bson:"size" validate:"required"`
	Material string             `bson:"metarial" validate:"required"`
}

// clothing schema input
type ClothingSchemaInput struct {
	Brand    string `bson:"brand" validate:"required"`
	Size     string `bson:"size" validate:"required"`
	Material string `bson:"metarial" validate:"required"`
}

// validate clothing input
func (c *ClothingSchemaInput) Validate() (bool, string) {
	// TODO: validate
	return true, ""
}

// ClothingProductRequest combines product and clothing data for API requests
type ClothingProductRequest struct {
	Product  ProductSchemaInput  `json:"product"`
	Clothing ClothingSchemaInput `json:"clothing"`
}

// validate clothing product request
func (c *ClothingProductRequest) Validate() (bool, string) {
	// TODO: validate
	return true, ""
}

// Electronic schema definition
type ElectronicSchema struct {
	ID           primitive.ObjectID `bson:"_id"`
	Manufacturer string             `bson:"manufacturer" validate:"required"`
	Model        string             `bson:"model" validate:"required"`
	Color        string             `bson:"color" validate:"required"`
}

// electronic schema input
type ElectronicSchemaInput struct {
	Manufacturer string `bson:"manufacturer" validate:"required"`
	Model        string `bson:"model" validate:"required"`
	Color        string `bson:"color" validate:"required"`
}

// validate electronic schema input
func (e *ElectronicSchemaInput) Validate() (bool, string) {
	// TODO: validate
	return true, ""
}

// ElectronicProductRequest combines product and electronic data for API requests
type ElectronicProductRequest struct {
	Product    ProductSchemaInput    `json:"product"`
	Electronic ElectronicSchemaInput `json:"electronic"`
}

// validate electronic product request
func (e *ElectronicProductRequest) Validate() bool {
	// TODO: validate
	return true
}
