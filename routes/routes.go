package routes

import (
	"html/template"
	"net/http"
	"sync"
    "github.com/alexedwards/scs/v2"

	"github.com/JorgeMG117/WizardsECommerce/middleware"
)

type Server struct {
	//Db          *sql.DB
    SessionManager *scs.SessionManager
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
	mux.HandleFunc("/", middleware.AuthenticationMiddleware(s.SessionManager, landingPage))
	mux.HandleFunc("/hello", s.Hello)
	mux.HandleFunc("/products", s.Products)

    // Cart
	mux.HandleFunc("/cart", cartPage)
	mux.HandleFunc("/get-cart-items", s.GetCart)
	mux.HandleFunc("/add-to-cart", s.AddToCart)
	mux.HandleFunc("/delete-from-cart", s.DeleteFromCart)

	mux.HandleFunc("/users", middleware.AuthenticationMiddleware(s.SessionManager, s.UsersPage))
	mux.HandleFunc("/getusers", s.GetUsersHandler)
	mux.HandleFunc("/login", s.LoginHandler)
	mux.HandleFunc("/register", s.RegisterHandler)
	//mux.HandleFunc("/logout", logoutHandler)
	return s.SessionManager.LoadAndSave(mux)
}
