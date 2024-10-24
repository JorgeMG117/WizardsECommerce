package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
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

var TemplateCache = make(map[string]*template.Template)

func LoadTemplates() error {
    //templates, err := filepath.Glob("views/templates/*.html")
    templates, err := filepath.Glob("views/prueba.html")
    if err != nil {
        return err
    }

    includes, err := filepath.Glob("views/includes/*.html")
    if err != nil {
        return err
    }

    baseTemplate := "views/base.html"

    fmt.Println(templates)
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

func (s *Server) Prueba(w http.ResponseWriter, r *http.Request) {
    tmpl, ok := TemplateCache["prueba.html"]
    if !ok {
        http.Error(w, "Could not load template", http.StatusInternalServerError)
        return
    }
    err := tmpl.ExecuteTemplate(w, "base", nil)
    if err != nil {
        http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
        return
    }
}

func ServeStatic() {
}

//TODO: Maybe add func RenderStatic

func (s *Server) Router() http.Handler {
	//th := timeHandler{format: "a"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.AuthenticationMiddleware(s.SessionManager, landingPage))

	mux.HandleFunc("/index", s.Index)
	mux.HandleFunc("/shop", s.Shop)
	mux.HandleFunc("/about", s.About)
	mux.HandleFunc("/prueba", s.Prueba)
	mux.HandleFunc("/contact", s.Contact)
	// mux.HandleFunc("/cart", s.Cart)

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


    // Serve static files from the 'static' directory
    fileServer := http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return s.SessionManager.LoadAndSave(mux)
}
