package models

import (
	"errors"
	"os"

	"github.com/JorgeMG117/WizardsECommerce/utils"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
}

var productFile string = "data/products.json"

func CreateProduct(product Product) error {
	var products []Product
	err := utils.ReadFile(productFile, &products)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	product.ID = len(products) + 1
	products = append(products, product)

	return utils.WriteFile(productFile, products)
}

func GetProducts() ([]Product, error) {
	var products []Product
	err := utils.ReadFile(productFile, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
