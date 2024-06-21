package routes

import (
	"html/template"
	"net/http"
	"sync"
)

type Server struct {
	//Db          *sql.DB
	mutex sync.Mutex
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/products.html"))
	tmpl.Execute(w, nil)
}

func (s *Server) Router() http.Handler {
	//th := timeHandler{format: "a"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", landingPage)
	mux.HandleFunc("/hello", s.Hello)
	mux.HandleFunc("/products", s.Products)
	return mux
}
