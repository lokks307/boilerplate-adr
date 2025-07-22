package responder

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/lokks307/adr-boilerplate/models"
	"github.com/lokks307/djson/v2"
)

func ResponderGetOneCustomer(
	ctx echo.Context,
	customer *models.Customer,
) error {
	res := djson.New().Put(djson.Object{
		"id":      customer.CustomerId,
		"name":    customer.FirstName + " " + customer.LastName,
		"company": customer.Company.String,
		"address": customer.Address.String,
		"city":    customer.City.String,
		"State":   customer.State.String,
		"phone":   customer.Phone.String,
	})

	return Response(ctx, http.StatusOK, res)
}
