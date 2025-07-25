package router

import (
	"github.com/lokks307/adr-boilerplate/action"
	"github.com/lokks307/adr-boilerplate/middleware"
)

var customerRoutes = []Route{
	{
		Method:     "POST",
		Path:       "/customers",
		Action:     action.InsertCustomer,
		Middleware: middleware.WithJwt,
	},
	{
		Method:     "GET",
		Path:       "/customers/:customer_id",
		Action:     action.GetCustomerInfo,
		Middleware: middleware.WithApiKey,
	},
}
