// models/cart.go
package models

type CartItem struct {
    Product  Product
    Quantity int
    Subtotal float64
}

type Cart map[int]*CartItem // Use pointer to CartItem for in-place updates

// AddItem adds a product to the cart or updates its quantity
func (c Cart) AddItem(product Product, quantity int) {
    if item, found := c[product.ID]; found {
        item.Quantity += quantity
    } else {
        c[product.ID] = &CartItem{
            Product:  product,
            Quantity: quantity,
            Subtotal: product.Price * float64(quantity),
        }
    }
    // Update subtotal
    c[product.ID].Subtotal = c[product.ID].Product.Price * float64(c[product.ID].Quantity)
}

// RemoveItem removes a product from the cart
func (c Cart) RemoveItem(productId int) {
    delete(c, productId)
}

// CalculateTotal calculates the total cost of the cart
func (c Cart) CalculateTotal() float64 {
    total := 0.0
    for _, item := range c {
        item.Subtotal = item.Product.Price * float64(item.Quantity)
        total += item.Subtotal
    }
    return total
}

// GetTotalItems returns the total number of items in the cart
func (c Cart) GetTotalItems() int {
    totalItems := 0
    for _, item := range c {
        totalItems += item.Quantity
    }
    return totalItems
}

