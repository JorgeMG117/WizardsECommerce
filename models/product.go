package models

import (
	"errors"
	"fmt"

    "gorm.io/gorm"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
}


func CreateProduct(db *gorm.DB, product *Product) error {
    return db.Create(product).Error
}

func GetProducts(db *gorm.DB) ([]Product, error) {
    var products []Product
    err := db.Find(&products).Error
    return products, err
}

func GetProductById(db *gorm.DB, id int) (*Product, error) {
    var product Product
    err := db.First(&product, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("product with ID %d not found", id)
        }
        return nil, err
    }
    return &product, nil
}

func GetFeaturedProducts(db *gorm.DB) ([]Product, error) {
    p, e := GetProducts(db)
    return p, e
}

