package action

import (
	"bytes"
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/lokks307/adr-boilerplate/types/e"
	"github.com/lokks307/djson/v2"
	"github.com/sirupsen/logrus"
)

func GetDjsonFromBody(ctx echo.Context, dvKey string) (*djson.JSON, error) {
	if ctx.Request().Body == nil {
		return nil, e.ActionErrGetDjsonFromBody1
	}

	bodyBytes, err := ioutil.ReadAll(ctx.Request().Body)
	ctx.Request().Body.Close()
	if err != nil {
		logrus.Error(err)
		return nil, e.ActionErrGetDjsonFromBody2
	}

	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	bodyJson := djson.New().Parse(string(bodyBytes))
	if err != nil {
		logrus.Error(err)
		return nil, e.ActionErrGetDjsonFromBody3
	}

	dv, ok := DjsonValidatorMap[dvKey]
	if dvKey == "" || !ok {
		return nil, e.ActionErrGetDjsonFromBody5
	}

	if !dv.IsValid(bodyJson) {
		return nil, e.ActionErrGetDjsonFromBody4
	}

	return bodyJson, nil
}
