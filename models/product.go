package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/JorgeMG117/WizardsECommerce/utils"
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

func GetProductById(id int) Product {
    products, err := GetProducts()
    utils.CheckError(err)

    for _, v := range products {
        if v.ID == id {
            return v
        }
    }
    
    fmt.Println("That product doesnt exist")

    return products[0] 
}
