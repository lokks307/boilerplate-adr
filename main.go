package main

import (
	"flag"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/lokks307/adr-boilerplate/domain"
	"github.com/lokks307/adr-boilerplate/env"
	"github.com/lokks307/adr-boilerplate/router"
)

func main() {
	config := flag.String("config", "dev-env.toml", "config filepath")
	portNum := flag.Int("port", 8080, "port number to open")
	debugFlag := flag.Bool("debug", false, "print out debug message")
	flag.Parse()

	if err := env.Setup(*config); err != nil {
		logrus.Error(err)
		return
	}

	echoRouter := router.InitRouter()
	if *debugFlag {
		logrus.SetLevel(logrus.TraceLevel)
	}

	if err := domain.DBLoad(); err != nil {
		logrus.Error("domain LoadDB err=", err)
		return
	}

	if err := echoRouter.Run(":" + strconv.Itoa(*portNum)); err != nil {
		logrus.Error("server running err=", err)
		return
	}
}
