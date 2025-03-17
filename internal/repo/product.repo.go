package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// create product interface
type IProductRepo interface {
	GetCollectoin() *mongo.Collection
	InsertProduct(interface{}) error
}

// variable interface
var vIProduct IProductRepo

// new IProductRepo
func NewProductRepo(v IProductRepo) {
    vIProduct = v
}

// get product interface
func GetProductRepo() IProductRepo {
	if (vIProduct == nil) {
		panic("implement product Repo not found for interface IProductRepo")
	}
    return vIProduct
}