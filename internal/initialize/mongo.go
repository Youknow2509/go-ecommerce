package initialize

import (
	"context"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

// init mongo
func InitMongo() {
	// initialize mongo connection
    initMongoConnection()
    // create database and collections
    createMongoCollections()
    // initialize mongo indexes
    initMongoIndexes()
    // initialize mongo migration
    initMongoMigration()
}

// init mongo connection
func initMongoConnection() {
	strConnection := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		global.Config.MongoDB.Username,
		global.Config.MongoDB.Password,
		global.Config.MongoDB.Host,
		global.Config.MongoDB.Port,
		global.Config.MongoDB.Database,
	)
    // establish connection with mongodb
	c, err := mongo.Connect(options.Client().ApplyURI(strConnection))
	if err != nil {
		panic(err)
	}
	// check the connection
	err = c.Ping(context.Background(), nil)
	if err != nil {
		global.Logger.Error("Error connecting to MongoDB", zap.Error(err))
		panic(err)
	}
	// set the connection to global
	global.MongoClient = c
}

// create mongo collections
func createMongoCollections() {
    // create collections product in mongodb
    err := global.MongoClient.Database(model.DatabaseProduct).CreateCollection(context.Background(), model.ProductCollection)
    if err != nil {
        // Handle error
		global.Logger.Error("Error creating product collection", zap.Error(err))
	}
    
    // create collections electronics
    err = global.MongoClient.Database(model.DatabaseProduct).CreateCollection(context.Background(), model.ElectronicCollection)
    if err != nil {
        // Handle error
		global.Logger.Error("Error creating electronic collection", zap.Error(err))
	}
    
    // create collections clothings
    err = global.MongoClient.Database(model.DatabaseProduct).CreateCollection(context.Background(), model.ClothingCollection)
    if err != nil {
        // Handle error
		global.Logger.Error("Error creating clothing collection", zap.Error(err))
	}

    // ...
}

// initialize mongo indexes
func initMongoIndexes() {
	// create index for product collection

}

// initialize mongo migration
func initMongoMigration() {
    // initialize mongo migration
	
}
