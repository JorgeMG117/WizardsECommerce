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
			<div class="container">
				<h1 class="mt-5">Your Cart</h1>
				{{if .}}
					<div class="row mt-4">
					{{range .}}
						<div class="col-md-4 mb-4">
							<div class="card">
								<img src="{{.ImageURL}}" class="card-img-top" alt="{{.Name}}">
								<div class="card-body">
									<h5 class="card-title">{{.Name}}</h5>
									<p class="card-text">{{.Description}}</p>
									<p class="card-text">$ {{.Price}}</p>
									<a href="#" class="btn btn-danger">Remove from Cart</a>
								</div>
							</div>
						</div>
					{{end}}
					</div>
				{{else}}
					<p>Your cart is empty.</p>
				{{end}}
			</div>
		`))

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
