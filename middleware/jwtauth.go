package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/lokks307/adr-boilerplate/auth"
	"github.com/lokks307/adr-boilerplate/responder"
	"github.com/lokks307/adr-boilerplate/types"
)

var WithJwt = []echo.MiddlewareFunc{JwtAuth}
var WithApiKey = []echo.MiddlewareFunc{CheckApiKey}

func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		logrus.Info("token string ", tokenString)
		tokenClaim, err := auth.ParseJWT(tokenString)
		if err != nil {
			return responder.Response(c, http.StatusUnauthorized, "token error!")
		}

		if tokenClaim.ExpireTimestamp < time.Now().Unix() {
			return responder.Response(c, http.StatusUnauthorized, "token expired!")
		}

		c.Set("customer_id", tokenClaim.CustomerID)

		return next(c)
	}
}

func CheckApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authString := c.Request().Header.Get("Authorization")

		if authString != types.API_KEY {
			return responder.Response(c, http.StatusUnauthorized, "Authorization error!")
		}

		return next(c)
	}
}
