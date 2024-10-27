package models

type CartItem struct {
    Product     Product
    Quantity    int
    Subtotal    float64
}

type Cart []CartItem 
