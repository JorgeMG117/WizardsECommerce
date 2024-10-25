package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
    "strconv"

	"github.com/JorgeMG117/WizardsECommerce/models"
	"github.com/JorgeMG117/WizardsECommerce/utils"
)

var cartMutex sync.Mutex
var cart []models.Product

func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
        cartMutex.Lock()
		defer cartMutex.Unlock()

		// Render the items in the cart, not the full product list
		tmpl := template.Must(template.New("cart-items").Parse(`
            <section id="cart" class="section-p1">
				{{if .}}
                    <table width="100%">
                        <thead>
                            <tr>
                                <td>Remove</td>
                                <td>Image</td>
                                <td>Product</td>
                                <td>Price</td>
                                <td>Quantity</td>
                                <td>Subtotal</td>
                            </tr>
                        </thead>
                        <tbody>
                            {{range .}}
                                <tr>
                                    <td><i class='bx bx-x-circle'></i></td>
                                    <td><img src="{{.ImageURL}}" alt="{{.Name}}"></td>
                                    <td>{{.Name}}</td>
                                    <td>{{.Price}}€</td>
                                    <td><input type="number" value="1"></td>
                                    <td>{{.Price}}€</td>
                                </tr>
                            {{end}}
                        </tbody>
				{{else}}
					<p>Your cart is empty.</p>
				{{end}}
            </section>
		`))
        //<a href="#" class="btn btn-danger" hx-post="/delete-from-cart" hx-target="closest .col-md-4" hx-swap="outerHTML swap:remove" hx-headers='{"Content-Type": "application/json"}' hx-vals='{"Id": {{.ID}}}'>Remove from Cart</a>

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, cart); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}


func (s *Server) AddToCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
        //var product models.Product
        r.ParseForm()
        productId := r.Form.Get("Id")

        // Get product by Id
        productIdInt, err := strconv.Atoi(productId)
        utils.CheckError(err)
        product := models.GetProductById(productIdInt)

        cart = append(cart, product) 

        // I get the Id, from there I can check the price and add it to the cart 
        // if I receive everything from the message somebody could send wrong price

        //fmt.Println("Content-Type Header:", r.Header.Get("Content-Type"))

        w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
        //var product models.Product
        r.ParseForm()
        productId := r.Form.Get("Id")

        // Get product by Id
        productIdInt, err := strconv.Atoi(productId)
        utils.CheckError(err)

        // ERROR
        for i := range cart {
            if cart[i].ID == productIdInt {
                cart[i] = cart[len(cart)-1]
                cart = cart[:len(cart)-1]
            }
        }

        w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
