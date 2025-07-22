package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/lokks307/adr-boilerplate/env"
	"github.com/lokks307/adr-boilerplate/middleware"
)

type Route struct {
	Method     string
	Path       string
	Action     echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
	DisableLog bool
}

type Router struct {
	Server *echo.Echo
	Routes []Route
}

func InitRouter() *Router {
	return &Router{
		Server: echo.New(),
		Routes: make([]Route, 0),
	}
}

func (s *Router) addRouter() {
	// FIXME: 각 router를 모두 append 해줘야 함
	s.Routes = append(s.Routes, baseRoutes...)
	s.Routes = append(s.Routes, customerRoutes...)
	// TODO:

	for _, r := range s.Routes {
		logrus.Infof("adding, [%s] %s", r.Method, r.Path)

		if r.DisableLog {
			env.DisableLogMap[r.Path] = true
		}

		s.Server.Add(
			r.Method,
			r.Path,
			r.Action,
			r.Middleware...)
	}
}

func (m *Router) Run(addr string) error {
	m.addRouter()

	// 미들웨어 선언
	m.Server.Use(middleware.NewDefaultEventLogger(), middleware.LogrusLoggerMiddleware())

	return m.Server.Start(addr)
}
