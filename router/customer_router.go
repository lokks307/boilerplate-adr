package router

import (
	"github.com/labstack/echo/v4"

	"github.com/lokks307/adr-boilerplate/action"
)

var customerRoutes = []Route{
	{
		Method:     "POST",
		Path:       "/customers",
		Action:     action.InsertCustomer,
		Middleware: []echo.MiddlewareFunc{},
	},
	{
		Method:     "GET",
		Path:       "/customers/:customer_id",
		Action:     action.GetCustomerInfo,
		Middleware: []echo.MiddlewareFunc{},
	},
}
