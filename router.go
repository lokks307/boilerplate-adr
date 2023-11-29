package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lokks307/adr-boilerplate/action/customer_action"
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

var RouteMap = []Route{
	{
		Method:     "GET",
		Path:       "/heartbeat",
		Action:     customer_action.EchoHello,
		Middleware: []echo.MiddlewareFunc{},
		DisableLog: true,
	},
	{
		Method:     "GET",
		Path:       "/customer/:customer_id",
		Action:     customer_action.GetCustomerInfo,
		Middleware: []echo.MiddlewareFunc{},
	},
}

type Router struct {
	Server *echo.Echo
}

func (m *Router) Init() {
	m.Server = echo.New()
	// TODO: setup global middleware

	for idx := range RouteMap {
		if RouteMap[idx].DisableLog {
			env.DisableLogMap[RouteMap[idx].Path] = true
		}

		m.Server.Add(RouteMap[idx].Method, RouteMap[idx].Path, RouteMap[idx].Action, RouteMap[idx].Middleware...)
	}

	m.Server.Use(middleware.LogrusLoggerMiddleware())
	// FIXME: 아래 미들웨어를 활성화하면 프로그램이 죽어요
	// m.Server.Use(middleware.NewDefaultEventLogger())
}

func (m *Router) Run(addr string) error {
	return m.Server.Start(addr)
}
