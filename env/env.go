package env

import (
	"path/filepath"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

var LOG_PATH = "/var/log/mediease.log"
var DisableLogMap map[string]bool
var Database DatabaseSetting

type DatabaseSetting struct {
	Type     string
	User     string
	Password string
	Host     string
	DBName   string
	SSLMode  string
	Port     int
}

func Setup(configPath string) error {
	rootDir, _ := filepath.Abs("./")

	var configToml *toml.Tree
	var err error
	if len(configPath) > 0 {
		configToml, err = toml.LoadFile(rootDir + configPath)
	} else {
		configToml, err = toml.LoadFile(rootDir + "\\config\\dev-env.toml")
	}
	if err != nil {
		logrus.Error(err)
		return err
	}

	dbToml := configToml.Get("db").(*toml.Tree)
	if err := dbToml.Unmarshal(&Database); err != nil {
		logrus.Error(err)
		return err
	}

	DisableLogMap = make(map[string]bool)

	return nil
}
