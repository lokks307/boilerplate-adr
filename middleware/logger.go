package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lokks307/adr-boilerplate/env"
	"github.com/lokks307/djson/v2"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func NewDefaultEventLogger() echo.MiddlewareFunc {
	fpLog, err := os.OpenFile(env.LOG_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fpLog,
	})
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "060102 15:04:05.000",
		FullTimestamp:   true,
		ForceColors:     true,
	})
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetLevel(logrus.InfoLevel)
}

func LogrusLoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			errFlag := false

			if err = next(c); err != nil {
				// TODO: handle error
				c.Error(err)
				errFlag = true
			}
			stop := time.Now()

			reqId := req.Header.Get(echo.HeaderXRequestID)
			if reqId == "" {
				reqId = res.Header().Get(echo.HeaderXRequestID)
			}

			logJson := djson.New()
			logJson.Put(djson.Object{
				"uri":      req.RequestURI,
				"req_host": c.RealIP(),
				"req_id":   reqId,
				"time":     stop.Sub(start).String(),
			})

			logJsonStr := logJson.String()
			logJsonStr = strings.TrimLeft(logJsonStr, "{")
			logJsonStr = strings.TrimRight(logJsonStr, "}")

			if errFlag || res.Status >= 400 {
				logrus.Warnf("[%s, %d] %s",
					req.Method,
					res.Status,
					logJsonStr,
				)
			} else {
				logrus.Infof("[%s, %d] %s",
					req.Method,
					res.Status,
					logJsonStr,
				)
			}
			return
		}
	}
}

func JsonLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			errFlag := false

			if err = next(c); err != nil {
				c.Error(err)
				errFlag = true
			}

			if req.RequestURI != "/heartbeat" {
				stop := time.Now()

				reqId := req.Header.Get(echo.HeaderXRequestID)
				if reqId == "" {
					reqId = res.Header().Get(echo.HeaderXRequestID)
				}

				fields := logrus.Fields{
					"request_id": reqId,
					"client_ip":  c.RealIP(),
					"latency":    stop.Sub(start).String(),
					"method":     req.Method,
					"uri":        req.RequestURI,
					"status":     res.Status,
					"referrer":   req.Host,
					"user_agent": req.UserAgent(),
				}

				entry := logrus.WithFields(fields)

				logrus.SetFormatter(&logrus.JSONFormatter{})
				if errFlag || res.Status >= 500 {
					entry.Error()
				} else if res.Status >= 400 {
					entry.Warn()
				}
			}

			return
		}
	}
}
