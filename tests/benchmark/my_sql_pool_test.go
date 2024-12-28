package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{
		Name: "test",
		// Other fields
	}

	if err := db.Create(&user).Error; err != nil {
		b.Fatalf("Error: %v", err)
	}
}

// go test -bench=. -benchmem

// func BenchmarkMaxOpenConns1(b *testing.B) {
// 	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := "root:root123@tcp(127.0.0.1:3306)/go_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error: %v", err)
// 	}

// 	// Check if the tables exist
// 	if !db.Migrator().HasTable(&User{}) {
// 		// drop table if exists
// 		if err := db.Migrator().DropTable(&User{}); err != nil {
// 			log.Fatalf("Error dropping table: %v", err)
// 		}
// 		// create table
// 		db.AutoMigrate(&User{})
// 		sqlDB, err := db.DB()
// 		if err != nil {
// 			log.Fatalf("Error get sqlDB from gorm.DB: %v", err)
// 		}
// 		sqlDB.SetMaxOpenConns(1)

// 		defer sqlDB.Close()

// 		b.RunParallel(func(p *testing.PB) {
// 			for p.Next() {
// 				insertRecord(b, db)
// 			}
// 		})
// 	}
// }

func BenchmarkMaxOpenConns10(b *testing.B) {

	dsn := "root:root123@tcp(127.0.0.1:3306)/go_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Check if the tables exist
	if !db.Migrator().HasTable(&User{}) {
		// drop table if exists
		if err := db.Migrator().DropTable(&User{}); err != nil {
			log.Fatalf("Error dropping table: %v", err)
		}
		// create table
		db.AutoMigrate(&User{})
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error get sqlDB from gorm.DB: %v", err)
		}
		sqlDB.SetMaxOpenConns(10)

		defer sqlDB.Close()

		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				insertRecord(b, db)
			}
		})
	}
}
