package routes

import (
	"fmt"
	"net/http"
    "strconv"
    "html/template"

	"github.com/JorgeMG117/WizardsECommerce/models"
)


func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
    // Get cart from session
    cartInterface := s.SessionManager.Get(r.Context(), "cart")
    var cart models.Cart
    if cartInterface != nil {
        cart, _ = cartInterface.(models.Cart)
    } else {
        cart = make(models.Cart)
    }

    total := cart.CalculateTotal()

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
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    r.ParseForm()
    productIdStr := r.Form.Get("Id")
    quantityStr := r.Form.Get("quantity")

    productId, err := strconv.Atoi(productIdStr)
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

    product, err := models.GetProductById(s.Db, productId)
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    // Get cart from session
    cartInterface := s.SessionManager.Get(r.Context(), "cart")
    var cart models.Cart
    if cartInterface != nil {
        cart, _ = cartInterface.(models.Cart)
    } else {
        cart = make(models.Cart)
    }

    // Use the AddItem method
    cart.AddItem(*product, quantity)

    // Save cart back to session
    s.SessionManager.Put(r.Context(), "cart", cart)

    // Calculate the total number of items in the cart
    cartItemCount := cart.GetTotalItems()

    fmt.Fprintf(w, `<span id="cart-count" class="cart-badge">%d</span>`, cartItemCount)
}


func (s *Server) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    r.ParseForm()
    productIdStr := r.Form.Get("Id")
    productId, err := strconv.Atoi(productIdStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    // Get cart from session
    cartInterface := s.SessionManager.Get(r.Context(), "cart")
    var cart models.Cart
    if cartInterface != nil {
        cart, _ = cartInterface.(models.Cart)
    } else {
        cart = make(models.Cart)
    }

    // Use the RemoveItem method
    cart.RemoveItem(productId)

    // Save updated cart to session
    s.SessionManager.Put(r.Context(), "cart", cart)

    // Recalculate total
    total := cart.CalculateTotal()

    data := struct {
        Items models.Cart
        Total float64
    }{
        Items: cart,
        Total: total,
    }

    // Render the updated cart section
    tmpl := template.Must(template.New("cart-section").ParseFiles("views/includes/cart-section.html"))
    w.Header().Set("Content-Type", "text/html")
    err = tmpl.Execute(w, data)
    if err != nil {
        fmt.Println("Error rendering template:", err)
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
    }
}


