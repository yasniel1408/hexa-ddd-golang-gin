package db

import (
	"fmt"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql/dao"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Migraciones automáticas
	errAutoMigrate := db.AutoMigrate(&dao.UserDao{})
	if errAutoMigrate != nil {
		fmt.Print(errAutoMigrate)
	}

	return db
}
