package db

import (
	"log"

	productEntities "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/entities"
	userEntities "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Migraciones automáticas
	db.AutoMigrate(&userEntities.User{})
	db.AutoMigrate(&productEntities.Product{})

	return db
}
