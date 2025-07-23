package action

import (
	"bytes"
	"io"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/lokks307/adr-boilerplate/e"
	"github.com/lokks307/adr-boilerplate/types"
	"github.com/lokks307/djson/v2"
)

func getDjsonFromBody(ctx echo.Context, dvKey string) (*djson.JSON, error) {
	if ctx.Request().Body == nil {
		return nil, e.ActionErrGetDjsonFromBody1
	}

	bodyBytes, err := io.ReadAll(ctx.Request().Body)
	ctx.Request().Body.Close()
	if err != nil {
		logrus.Error(err)
		return nil, e.ActionErrGetDjsonFromBody2
	}
	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	bodyJson := djson.New().Parse(string(bodyBytes))

	dv, ok := djsonValidatorMap[dvKey]
	if dvKey == "" || !ok {
		return nil, e.ActionErrGetDjsonFromBody3
	}

	if !dv.IsValid(bodyJson) {
		return nil, e.ActionErrGetDjsonFromBody4
	}

	return bodyJson, nil
}

func GetIntPathParamToMap(
	ctx echo.Context,
	keys ...string,
) (map[string]int64, error) {
	res := make(map[string]int64)

	for _, eachKey := range keys {
		if eachKey == types.DEFAULT_EMPTY_STRING {
			continue
		}

		id, err := strconv.ParseInt(ctx.Param(eachKey), 10, 64)
		if err != nil {
			return nil, e.ErrorWrap(err, eachKey)
		}

		if id <= types.INVALID_ID {
			return nil, e.ErrorWrap(e.ActionErrInvalidPathParam, "params: (", eachKey, ")=", id)
		}

		res[eachKey] = id
	}

	return res, nil
}
