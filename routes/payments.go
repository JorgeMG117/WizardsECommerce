package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JorgeMG117/WizardsECommerce/models"
	"github.com/JorgeMG117/WizardsECommerce/utils"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)


func (s *Server) CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        fmt.Println("Error parsing form:", err)
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }
    productIds_str := r.Form["productIds"]
    quantities_str := r.Form["quantities"]

    var lineItems []*stripe.CheckoutSessionLineItemParams

    productIds, err := utils.ConvertStringsToInts(productIds_str)
    utils.CheckError(err)
    quantities, err := utils.ConvertStringsToInts(quantities_str)
    utils.CheckError(err)
    products, err := models.GetProductsByIds(s.Db, productIds)

    for i := range products {
        p := &stripe.CheckoutSessionLineItemParams{
            PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
                Currency: stripe.String("eur"),
                ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams {
                    Name: stripe.String(products[i].Name),
                },
                UnitAmount: stripe.Int64(int64(products[i].Price * 100)),
            },
            Quantity: stripe.Int64(int64(quantities[i])),
        }
        lineItems = append(lineItems, p)
    }
                            
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
