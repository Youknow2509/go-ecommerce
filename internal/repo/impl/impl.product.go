package impl

import (
	"context"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// clothing impl implementation
type ClothingRepo struct {
	mongoClient *mongo.Client
}

// InsertProduct implements repo.IProductRepo.
func (c *ClothingRepo) InsertProduct(r interface{}) error {
	input := r.(model.ClothingProductRequest)
	objID := primitive.NewObjectID()
	// create schema input product
	schema_product := model.ProductSchema{
		Product_ID:          objID,
		Product_Name:        input.Product.Product_Name,
		Product_Thumb:       input.Product.Product_Thumb,
		Product_Description: input.Product.Product_Description,
		Product_Price:       input.Product.Product_Price,
		Product_Quantity:    input.Product.Product_Quantity,
		Product_Type:        model.Product_Type_Clothing,
		Product_Shop:        input.Product.Product_Shop,
		Product_Atributes:   input.Clothing,
		Product_Created_At:  time.Now(),
		Product_Updated_At:  time.Now(),
	}
	// insert product
	err := NewProductRepo(c.mongoClient).InsertProduct(schema_product)
	if err != nil {
		global.Logger.Error("Error inserting product", zap.Error(err))
		return err
	}
	// create schema clothing
	schema_clothing := model.ClothingSchema{
		ID:       objID,
		Brand:    input.Clothing.Brand,
		Size:     input.Clothing.Size,
		Material: input.Clothing.Material,
	}
	res, err := c.GetCollectoin().InsertOne(
		context.Background(),
		schema_clothing,
	)
	if err != nil {
		global.Logger.Error("Error inserting product", zap.Error(err))
		return err
	}
	global.Logger.Info("Product inserted", zap.Any("product_id", res.InsertedID))
	return nil
}

// GetCollectoin implements repo.IProductRepo.
func (c *ClothingRepo) GetCollectoin() *mongo.Collection {
	return c.mongoClient.Database(model.DatabaseProduct).Collection(model.ClothingCollection)
}

// new clothing impl
func NewClothingRepo(mc *mongo.Client) repo.IProductRepo {
	return &ClothingRepo{
		mongoClient: mc,
	}
}

// product implements repo.IProductRepo
type ProductRepo struct {
	mongoClient *mongo.Client
}

// GetCollectoin implements repo.IProductRepo.
func (p *ProductRepo) GetCollectoin() *mongo.Collection {
	return p.mongoClient.Database(model.DatabaseProduct).Collection(model.ProductCollection)
}

// InsertProduct implements repo.IProductRepo.
func (p *ProductRepo) InsertProduct(r interface{}) error {
	input := r.(model.ProductSchema)
	// insert product
	res, err := p.GetCollectoin().InsertOne(
		context.Background(),
		input,
	)
	if err != nil {
		global.Logger.Error("Error inserting product", zap.Error(err))
		return err
	}
	global.Logger.Info("Product inserted", zap.Any("product_id", res.InsertedID))
	return nil
}

// new product repo
func NewProductRepo(mc *mongo.Client) repo.IProductRepo {
	return &ProductRepo{
		mongoClient: mc,
	}
}
