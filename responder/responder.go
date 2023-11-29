package responder

import (
	"github.com/labstack/echo/v4"
	"github.com/lokks307/djson/v2"
)

var JsonResponder *JSONResponder
var TxtResponder *TextResponder

type JSONResponder struct {
}

func NewJSONResponder() *JSONResponder {
	return &JSONResponder{}
}

func (m *JSONResponder) Response(ctx echo.Context, code int, val interface{}) error {
	// FIXME: default type handling??
	switch t := val.(type) {
	case *djson.JSON:
		return ctx.String(code, t.String())
	default:
		return ctx.JSON(code, t)
	}
}

type TextResponder struct {
}

func NewTextResponder() *TextResponder {
	return &TextResponder{}
}

func (m *TextResponder) Response(ctx echo.Context, code int, val string) error {
	return ctx.String(code, val)
}

func init() {
	JsonResponder = NewJSONResponder()
	TxtResponder = NewTextResponder()
}
