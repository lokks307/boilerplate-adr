package env

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"

	"github.com/lokks307/adr-boilerplate/e"
)

var LOG_PATH = "/var/log/boilerplate.log"
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

var TZ *time.Location
var IsProd bool

func Setup(configPath string) error {
	var terr error
	TZ, terr = time.LoadLocation("Asia/Seoul")
	if terr != nil {
		return terr
	}

	if strings.Contains(configPath, "prod") {
		IsProd = true
	} else {
		IsProd = false
	}

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
		return e.ErrorWrap(err)
	}

	dbToml := configToml.Get("db").(*toml.Tree)
	if err := dbToml.Unmarshal(&Database); err != nil {
		logrus.Error(err)
		return e.ErrorWrap(err)
	}

	DisableLogMap = make(map[string]bool)

	return nil
}
