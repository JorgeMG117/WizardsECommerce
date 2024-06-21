package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/JorgeMG117/WizardsECommerce/models"
)

func (s *Server) Products(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		s.mutex.Lock()

		products, _ := models.GetProducts()

		s.mutex.Unlock()
		// TODO: Create checkError function
		// if err != nil {

		// }

		tmpl := template.Must(template.New("products").Parse(`
			{{range .}}
				<div class="col-md-4 mb-4">
					<div class="card">
						<img src="{{.ImageURL}}" class="card-img-top" alt="{{.Name}}">
						<div class="card-body">
							<h5 class="card-title">{{.Name}}</h5>
							<p class="card-text">{{.Description}}</p>
							<p class="card-text">$ {{.Price}}</p>
							<a href="#" class="btn btn-primary">Add to Cart</a>
						</div>
					</div>
				</div>
			{{end}}
		`))

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, products)
	case "POST":
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("POST - Products\n"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
