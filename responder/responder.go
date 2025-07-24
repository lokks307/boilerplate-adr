package responder

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/lokks307/adr-boilerplate/env"
	"github.com/lokks307/adr-boilerplate/types"
	"github.com/lokks307/djson/v2"
)

func Response(c echo.Context, code int, res interface{}) error {
	switch t := res.(type) {
	case *djson.JSON:
		return c.String(code, t.String())

	case string:
		return c.String(code, res.(string))

	case error:
		logrus.Error("error response should use ResponseError()!")
		SendTeamsMsg(fmt.Sprintf("[%s] %s", c.Request().Method, c.Path()),
			"Response 구현 실수!",
			"ResponseError 써야함!")
		return c.String(code, t.Error())

	default:
		return c.String(code, "")
	}
}
func ResponseError(c echo.Context, code int, res error, internalMsg any) error {
	var internalMsgStr string

	switch v := internalMsg.(type) {
	case string:
		internalMsgStr = v
	case error:
		internalMsgStr = v.Error()
	default:
		internalMsgStr = "내부 메시지 형태 에러!"
	}

	logrus.Error(internalMsg) // 콘솔 출력
	SendTeamsMsg(             // teams 전송
		fmt.Sprintf("[%s] %s", c.Request().Method, c.Path()),
		fmt.Sprintf("%d error!", code),
		internalMsgStr)

	return c.String(code, res.Error())
}

func ResponseEmptyObjectOK(ctx echo.Context) error {
	return Response(ctx, http.StatusOK, djson.NewObject())
}

func SendTeamsMsg(path, title string, errMsg string) {
	if !env.IsProd {
		path = "[DEVELOP] " + path
	}

	payloadBytes := `{
		"@type": "MessageCard",
		"@context": "http://schema.org/extensions",
		"themeColor": "0076D7",
		"summary": "AlimTalk API server error",
		"sections": [{
			"activityTitle": "` + path + `",
			"activitySubTitle": "AlimTalk API: ` + title + `",
			"facts": [{
				"name": "message",
				"value": "` + jsonEscapeString(errMsg) + `"
			}],
			"markdown": true
		}]
	}`

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", types.TeamsWebhookUrl, bytes.NewBuffer([]byte(payloadBytes)))
	if err != nil {
		fmt.Printf("Failed to create HTTP request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Failed to send POST request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("POST request failed with status code: %d", resp.StatusCode)
		return
	}
}

func jsonEscapeString(value string) string {
	var sb strings.Builder
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch c {
		case '\b', '\f', '\r', '\n', '\t':
			sb.WriteByte(' ')
		case '"', '\\':
			sb.WriteByte('\\')
			sb.WriteByte(c)
			sb.WriteByte(c)
		default:
			sb.WriteByte(c)
		}
	}

	return sb.String()
}
