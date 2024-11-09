package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
    
	//"github.com/JorgeMG117/WizardsECommerce/models"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)


func (s *Server) CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
    fmt.Println("En payments")
    type product struct {
        priceInCents int64
        name string
    }

    storeItems := make(map[int]product)
    storeItems[1] = product{ priceInCents: 1000, name: "Learn React Today" }
    storeItems[2] = product{ priceInCents: 2000, name: "Learn Go Today" }

    err := r.ParseForm()
    if err != nil {
        fmt.Println("Error parsing form:", err)
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }
    productIds := r.Form["productIds"]
    quantities := r.Form["quantities"]

    fmt.Println("Product IDs:", productIds)
    fmt.Println("Quantities:", quantities)
    var lineItems []*stripe.CheckoutSessionLineItemParams
    for range productIds {
        fmt.Println("Adding products")
        // From every id get the product
        p := &stripe.CheckoutSessionLineItemParams{
            PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
                Currency: stripe.String("eur"),
                ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams {
                    Name: stripe.String(storeItems[1].name),
                },
                UnitAmount: stripe.Int64(storeItems[1].priceInCents),
            },
            Quantity: stripe.Int64(1),
        }
        lineItems = append(lineItems, p)
    }
                            
	//products, _ := models.GetProducts()
    params := &stripe.CheckoutSessionParams{
    LineItems: lineItems,
    Mode: stripe.String("payment"),
    SuccessURL: stripe.String(os.Getenv("SERVER_URL")+ "/success.html"),
    CancelURL: stripe.String(os.Getenv("SERVER_URL")+ "/cancel.html"),
    }

    se, err := session.New(params)

    if err != nil {
        log.Printf("session.New: %v", err)
    }

    http.Redirect(w, r, se.URL, http.StatusSeeOther)
}
