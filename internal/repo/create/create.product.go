package create

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/repo"
	"github.com/Youknow2509/go-ecommerce/internal/repo/impl"
)

// const enum create product repo - save in model store
func CreateProduct(name string) repo.IProductRepo {
	switch name {
	case model.ClothingCollection:
		return impl.NewClothingRepo(global.MongoClient)
	case model.ElectronicCollection:
		return nil // TODO
	default:
		panic("Invalid product collection")
	}
}
