package action

import (
	"errors"

	"github.com/lokks307/adr-boilerplate/e"
	"github.com/lokks307/djson/v2"
)

const (
	INSERT_CUSTOMER = "INSERT_CUSTOMER"
)

var djsonValidatorMap = make(map[string]*djson.Validator)
var djsonValidatorSyntax = map[string]string{
	INSERT_CUSTOMER: `{
		"type": "OBJECT",
		"object": {
			"first_name": {
				"type": "STRING",
				"required": true
			},
			"last_name": {
				"type": "STRING",
				"required": true
			},
			"email": {
				"type": "STRING",
				"required": true
			}
		}
	}`,
}

func InitValidator() error {
	for k, v := range djsonValidatorSyntax {
		dv := djson.NewValidator()
		if dv.Compile(v) {
			djsonValidatorMap[k] = dv
		} else {
			return e.ErrorWrap(e.ErrInitValidator, errors.New(k))
		}
	}
	return nil
}
