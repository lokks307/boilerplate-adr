package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/lokks307/adr-boilerplate/responder"
)

var baseRoutes = []Route{
	{
		Method:     "GET",
		Path:       "/heartbeat",
		Action:     Heartbeat,
		Middleware: nil,
	},
}

func Heartbeat(ctx echo.Context) error {
	return responder.Response(ctx, http.StatusOK, "ADR API OK")
}
