package configs

import (
    "log"

    "github.com/JorgeMG117/WizardsECommerce/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("data/ecommerce.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrate the schema
    if err := db.AutoMigrate(&models.Product{}, &models.User{}); err != nil {
        log.Fatalf("Failed to migrate database schema: %v", err)
    }

    return db
}

