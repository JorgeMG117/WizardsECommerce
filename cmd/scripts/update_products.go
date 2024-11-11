package main

import (
    "encoding/json"
    "fmt"
    "os"
    "log"
    "path/filepath"

    "github.com/JorgeMG117/WizardsECommerce/configs"
    "github.com/JorgeMG117/WizardsECommerce/models"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

func main() {
    // Connect to the database
    db := configs.ConnectDB()

    // Load products from JSON file
    products, err := loadProductsFromJSON("data/products.json")
    if err != nil {
        log.Fatalf("Failed to load products: %v", err)
    }

    // Update the database
    if err := updateProductsInDB(db, products); err != nil {
        log.Fatalf("Failed to update products in database: %v", err)
    }

    fmt.Println("Products have been successfully updated in the database.")
}

func loadProductsFromJSON(filename string) ([]models.Product, error) {
    absPath, err := filepath.Abs(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to get absolute path: %w", err)
    }

    data, err := os.ReadFile(absPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read file %s: %w", absPath, err)
    }

    var products []models.Product
    err = json.Unmarshal(data, &products)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    return products, nil
}

func updateProductsInDB(db *gorm.DB, products []models.Product) error {
    for _, product := range products {
        // Use GORM's Upsert feature
        err := db.Clauses(clause.OnConflict{
            UpdateAll: true,
        }).Create(&product).Error
        if err != nil {
            return fmt.Errorf("failed to upsert product ID %d: %w", product.ID, err)
        }
    }
    return nil
}

