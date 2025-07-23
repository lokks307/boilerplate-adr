package action

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/lokks307/adr-boilerplate/domain/biz"
	"github.com/lokks307/adr-boilerplate/e"
	"github.com/lokks307/adr-boilerplate/responder"
)

func InsertCustomer(ctx echo.Context) error {
	bodyJson, err := getDjsonFromBody(ctx, INSERT_CUSTOMER)
	if err != nil {
		return responder.ResponseError(ctx, http.StatusBadRequest, e.ActionErrInvalidBody, e.ErrorWrap(err))
	}

	firstName := bodyJson.String("first_name")
	lastName := bodyJson.String("last_name")
	email := bodyJson.String("email")

	if err := biz.Customer().InsertCustomer(firstName, lastName, email); err != nil {
		return responder.ResponseError(ctx, http.StatusInternalServerError, e.ActionErrInternalServer, e.ErrorWrap(err))
	}

	return responder.ResponseEmptyObjectOK(ctx)
}

func GetCustomerInfo(ctx echo.Context) error {
	idMap, err := GetIntPathParamToMap(ctx, "customer_id")
	if err != nil {
		return responder.ResponseError(ctx, http.StatusBadRequest, e.ActionErrInvalidPathParam, e.ErrorWrap(err))
	}
	cid := idMap["customer_id"]

	customer, err := biz.Customer().ReadCustomerByID(cid)
	if err != nil {
		return responder.ResponseError(ctx, http.StatusNotFound, e.ActionErrNotFound, e.ErrorWrap(err))
	}

	return responder.ResponderGetOneCustomer(ctx, customer)
}
