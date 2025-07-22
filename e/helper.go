package e

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func buildLogString(fileName string, line int, ok bool, msgIn []interface{}) string {
	msgList := make([]string, 0)

	for _, msg := range msgIn {
		switch v := msg.(type) {
		case error:
			msgList = append(msgList, v.Error())
		case []string:
			msgList = append(msgList, fmt.Sprintf("[%s]", strings.Join(v, ",")))
		case []interface{}:
			vx := make([]string, 0)
			for _, xx := range v {
				vx = append(vx, fmt.Sprintf("%v", xx))
			}
			msgList = append(msgList, fmt.Sprintf("[%s]", strings.Join(vx, ",")))
		default:
			msgList = append(msgList, fmt.Sprintf("%v", v))
		}
	}

	if ok {
		return fmt.Sprintf("%s (%s:%d)", strings.Join(msgList, " "), filepath.Base(fileName), line)
	}

	return strings.Join(msgList, " ")
}

func ErrorWrap(err error, msgIn ...interface{}) error {
	_, filename, line, ok := runtime.Caller(1)
	msg := buildLogString(filename, line, ok, msgIn)

	return errors.Join(err, errors.New(" >> "+msg))
}
