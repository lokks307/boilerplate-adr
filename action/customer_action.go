package action

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/lokks307/adr-boilerplate/domain/usecase"
	"github.com/lokks307/adr-boilerplate/e"
	"github.com/lokks307/adr-boilerplate/responder"
)

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

	return responder.ResponderGetOneCustomer(ctx, customer)
}
