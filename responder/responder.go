package responder

import (
	"github.com/labstack/echo/v4"

	"github.com/lokks307/djson/v2"
)

func Response(ctx echo.Context, code int, val interface{}) error {
	switch t := val.(type) {
	case *djson.JSON:
		return ctx.String(code, t.String())
	case error:
		return ctx.String(code, t.Error())
	default:
		// FIXME: default type handling??
		return ctx.JSON(code, t)
	}
}
