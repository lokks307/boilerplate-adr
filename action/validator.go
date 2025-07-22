package action

import "github.com/lokks307/djson/v2"

const (
	INSERT_VALUE = "INSERT_VALUE"
)

var DjsonValidatorMap map[string]*djson.Validator
var DjsonValidatorSyntax = map[string]string{
	INSERT_VALUE: `{
		"type": "OBJECT",
		"object": {
			"key": {
				"type": "STRING",
				"required": true
			},
			"value": {
				"type": "STRING",
				"required": true
			}
		}
	}`,
}
