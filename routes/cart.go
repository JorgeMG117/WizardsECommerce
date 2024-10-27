package routes

import (
	"fmt"
	"net/http"
	"sync"
    "strconv"
    "html/template"

	"github.com/JorgeMG117/WizardsECommerce/models"
)

var cartMutex sync.Mutex
var cart models.Cart



func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
    cartMutex.Lock()
    defer cartMutex.Unlock()

    // Get cart from session
    /*cartInterface := s.SessionManager.Get(r.Context(), "cart")
    var cart []models.CartItem
    if cartInterface != nil {
        cart = cartInterface.([]models.CartItem)
    }
    */

    total := 0.0
    for i := range cart {
        cart[i].Subtotal = cart[i].Product.Price * float64(cart[i].Quantity)
        total += cart[i].Subtotal
    }

    data := struct {
        Items models.Cart
        Total float64
    }{
        Items: cart,
        Total: total,
    }

    s.RenderTemplate(w, "cart.html", data)
}



func (s *Server) AddToCart(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)
    switch r.Method {
    case "POST":
        r.ParseForm()
        productId := r.Form.Get("Id")
        quantityStr := r.Form.Get("quantity")
    
        productIdInt, err := strconv.Atoi(productId)
        if err != nil {
            http.Error(w, "Invalid product ID", http.StatusBadRequest)
            return
        }
    
        quantity := 1
        if quantityStr != "" {
            quantityInt, err := strconv.Atoi(quantityStr)
            if err == nil && quantityInt > 0 {
                quantity = quantityInt
            }
        }
    
        product := models.GetProductById(productIdInt)
        /*TODO
        if product == nil {
            http.Error(w, "Product not found", http.StatusNotFound)
            return
        }
        */
    
        // Get cart from session
        /*
        cartInterface := s.SessionManager.Get(r.Context(), "cart")
        var cart []models.CartItem
        if cartInterface != nil {
            cart = cartInterface.([]models.CartItem)
        }
        */
    
        // Add or update cart item
        found := false
        for i := range cart {
            if cart[i].Product.ID == productIdInt {
                cart[i].Quantity += quantity
                found = true
                break
            }
        }
        if !found {
            cartItem := models.CartItem{
                Product:  product,
                Quantity: quantity,
            }
            cart = append(cart, cartItem)
        }
    
        // Save cart back to session
        //s.SessionManager.Put(r.Context(), "cart", cart)
    
        w.WriteHeader(http.StatusOK)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}



func (s *Server) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)
    switch r.Method {
    case "POST":
        r.ParseForm()
        productId := r.Form.Get("Id")
        productIdInt, err := strconv.Atoi(productId)
        if err != nil {
            http.Error(w, "Invalid product ID", http.StatusBadRequest)
            return
        }

        cartMutex.Lock()
        defer cartMutex.Unlock()

        // Get cart from session
        /*
        cartInterface := s.SessionManager.Get(r.Context(), "cart")
        var cart []models.CartItem
        if cartInterface != nil {
            cart = cartInterface.([]models.CartItem)
        }
        */

        // Remove item
        for i := 0; i < len(cart); i++ {
            if cart[i].Product.ID == productIdInt {
                cart = append(cart[:i], cart[i+1:]...)
                break
            }
        }
        fmt.Println(cart)

        // Save updated cart to session
        //s.SessionManager.Put(r.Context(), "cart", cart)

        // Recalculate total
        total := 0.0
        for i := range cart {
            cart[i].Subtotal = cart[i].Product.Price * float64(cart[i].Quantity)
            total += cart[i].Subtotal
        }

        data := struct {
            Items models.Cart
            Total float64
        }{
            Items: cart,
            Total: total,
        }

        // Render the updated cart section
        tmpl := template.Must(template.New("cart-section").ParseFiles("views/cart-section.html"))
        w.Header().Set("Content-Type", "text/html")
        err = tmpl.Execute(w, data)
        if err != nil {
            fmt.Println("Error rendering template:", err)
            http.Error(w, "Error rendering template", http.StatusInternalServerError)
        }
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

