package routes

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/JorgeMG117/WizardsECommerce/middleware"
)

type Server struct {
	//Db          *sql.DB
	mutex sync.Mutex
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/products.html"))
	tmpl.Execute(w, nil)
}

func cartPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/cart.html"))
	tmpl.Execute(w, nil)

}

func (s *Server) Router() http.Handler {
	//th := timeHandler{format: "a"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", landingPage)
	mux.HandleFunc("/hello", s.Hello)
	mux.HandleFunc("/products", s.Products)

    // Cart
	mux.HandleFunc("/cart", cartPage)
	mux.HandleFunc("/get-cart-items", s.GetCart)
	mux.HandleFunc("/add-to-cart", s.AddToCart)


	mux.HandleFunc("/users", middleware.AuthenticationMiddleware(s.UsersPage))
	mux.HandleFunc("/getusers", s.GetUsersHandler)
	mux.HandleFunc("/login", s.LoginHandler)
	//mux.HandleFunc("/logout", logoutHandler)
	return mux
}
