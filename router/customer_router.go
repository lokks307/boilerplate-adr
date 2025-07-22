package router

import (
	"github.com/labstack/echo/v4"

	"github.com/lokks307/adr-boilerplate/action"
)

var customerRoutes = []Route{
	{
		Method:     "GET",
		Path:       "/customer/:customer_id",
		Action:     action.GetCustomerInfo,
		Middleware: []echo.MiddlewareFunc{},
	},
}
