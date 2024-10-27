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

		tmpl := template.Must(template.New("products").Parse(`
			{{range .}}
                <div class="pro">
                    <img class="shirt" src="{{.ImageURL}}" alt="{{.Name}}">
                    <div class="des">
                        <span>{{.Brand}}</span>
                        <h5>{{.Name}}</h5>
                        <div class="star">
                            <i class='bx bxs-star' ></i>
                            <i class='bx bxs-star' ></i>
                            <i class='bx bxs-star' ></i>
                            <i class='bx bxs-star' ></i>
                            <i class='bx bxs-star' ></i>
                        </div>
                        <h4>₹{{.Price}}</h4>
                    </div>
                    <a hx-post="/add-to-cart" hx-swap="none" hx-headers='{"Content-Type": "application/json"}' hx-vals='{"Id": {{.ID}}}'><i class='bx bx-cart cart'></i></a>
                </div>
			{{end}}
		`))

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, products)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

