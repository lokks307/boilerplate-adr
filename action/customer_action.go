package action

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/lokks307/adr-boilerplate/domain/usecase"
	"github.com/lokks307/adr-boilerplate/e"
	"github.com/lokks307/adr-boilerplate/responder"
	"github.com/lokks307/djson/v2"
)

func EchoHello(ctx echo.Context) error {
	return responder.Response(ctx, http.StatusOK, "Hello ADR!")
}

func GetCustomerInfo(ctx echo.Context) error {
	idMap, err := GetIntPathParamToMap(ctx, "customer_id")
	if err != nil {
		return responder.ResponseError(ctx, http.StatusBadRequest, e.ActionErrInvalidPathParam, e.ErrorWrap(err))
	}

	cid := idMap["customer_id"]

	customer, err := usecase.Customer().ReadCustomerByID(cid)
	if err != nil {
		return responder.ResponseError(ctx, http.StatusNotFound, e.ActionErrNotFound, e.ErrorWrap(err))
	}

	res := djson.New().Put(djson.Object{
		"id":      customer.CustomerId,
		"name":    customer.FirstName + " " + customer.LastName,
		"company": customer.Company.String,
		"address": customer.Address.String,
		"city":    customer.City.String,
		"State":   customer.State.String,
		"phone":   customer.Phone.String,
	})

	return responder.Response(ctx, http.StatusOK, res)
}
