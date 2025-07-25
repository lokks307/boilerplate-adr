package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/lokks307/adr-boilerplate/e"
	"github.com/lokks307/adr-boilerplate/types"
)

type TokenClaims struct {
	jwt.RegisteredClaims

	CustomerID      int64 `json:"customer_id"`
	ExpireTimestamp int64
}

func GenJWT(customerID int64) (string, error) {

	tokenClaims := TokenClaims{
		CustomerID:      customerID,
		ExpireTimestamp: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenStr, err := newToken.SignedString([]byte(types.JWT_KEY))
	if err != nil {
		return "", e.ErrorWrap(err)
	}

	return tokenStr, nil
}

func ParseJWT(tokenString string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(types.JWT_KEY), nil
		},
	)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		} else {
			return nil, errors.New("TokenClaims Not Valid")
		}
	}

	return nil, err
}
