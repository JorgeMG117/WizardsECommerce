package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/alexedwards/scs/v2"

	"github.com/JorgeMG117/WizardsECommerce/middleware"
	"github.com/JorgeMG117/WizardsECommerce/models"
)

type Server struct {
	//Db          *sql.DB
    SessionManager *scs.SessionManager
	mutex sync.Mutex
}


var TemplateCache = make(map[string]*template.Template)

func LoadTemplates() error {
    templates, err := filepath.Glob("views/templates/*.html")
    if err != nil {
        return err
    }

    includes, err := filepath.Glob("views/includes/*.html")
    if err != nil {
        return err
    }

    baseTemplate := "views/base.html"

    for _, tmplFile := range templates {
        // Combine base, includes, and the current template file
        files := append([]string{baseTemplate}, includes...)
        files = append(files, tmplFile)

        tmpl, err := template.ParseFiles(files...)
        if err != nil {
            return err
        }

        // Use the base filename as the key in the cache
        tmplName := filepath.Base(tmplFile)
        fmt.Println("Loading template:", tmplName)
        TemplateCache[tmplName] = tmpl
    }

    return nil
}


type TemplateData struct {
    CartItemCount int
    Data          interface{}
}

func (s *Server) RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {

    tmpl, ok := TemplateCache[templateName]
    if !ok {
        http.Error(w, "Could not load template", http.StatusInternalServerError)
        return
    }

    cartItemCount := 1
    td := TemplateData{
        CartItemCount: cartItemCount,
        Data:          data,
    }

    err := tmpl.ExecuteTemplate(w, "base", td)
    if err != nil {
        http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
        return
    }
}


func (s *Server) Router() http.Handler {
	//th := timeHandler{format: "a"}
	mux := http.NewServeMux()
	//mux.HandleFunc("/", middleware.AuthenticationMiddleware(s.SessionManager, landingPage))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        s.RenderTemplate(w, "index.html", nil)
    })
	mux.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		s.mutex.Lock()
		products, _ := models.GetProducts()
		s.mutex.Unlock()
        s.RenderTemplate(w, "shop.html", products)
    })
	mux.HandleFunc("/cart", s.GetCart)
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        s.RenderTemplate(w, "contact.html", nil)
    })
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        s.RenderTemplate(w, "about.html", nil)
    })
	mux.HandleFunc("/product/", s.Product)

	// mux.HandleFunc("/hello", s.Hello)
	mux.HandleFunc("/products", s.Products)

    // Cart
	// mux.HandleFunc("/cart", cartPage)
	// mux.HandleFunc("/get-cart-items", s.GetCart)
	mux.HandleFunc("/add-to-cart", s.AddToCart)
	mux.HandleFunc("/delete-from-cart", s.DeleteFromCart)

	mux.HandleFunc("/users", middleware.AuthenticationMiddleware(s.SessionManager, s.UsersPage))
	mux.HandleFunc("/getusers", s.GetUsersHandler)

	mux.HandleFunc("/login", s.LoginHandler)
	mux.HandleFunc("/register", s.RegisterHandler)
	//mux.HandleFunc("/logout", logoutHandler)


    // Stripe
	mux.HandleFunc("/create-checkout-session", middleware.AuthenticationMiddleware(s.SessionManager, s.CreateCheckoutSession))



    // Serve static files from the 'static' directory
    fileServer := http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return s.SessionManager.LoadAndSave(mux)
}
