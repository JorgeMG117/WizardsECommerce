package routes

import(
    "fmt"
    "net/http"
    "strings"
    "strconv"

	"github.com/JorgeMG117/WizardsECommerce/models"
)

func (s *Server) Product(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
        idStr := strings.TrimPrefix(r.URL.Path, "/product/")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.NotFound(w, r)
            return
        }

        // Fetch product details based on ID
        product, err := models.GetProductById(s.Db, id)
        fmt.Println(product)
        // http.NotFound(w, r)

        featuredProducts, err := models.GetFeaturedProducts(s.Db)
        if err != nil {
            http.NotFound(w, r)
            return
        }

        data := struct {
            Product          *models.Product
            FeaturedProducts []models.Product
        }{
            Product:          product,
            FeaturedProducts: featuredProducts,
        }

        s.RenderTemplate(w, "product.html", data)


	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

