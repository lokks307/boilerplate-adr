package domain

import (
	"context"
	"database/sql"

	"github.com/lokks307/adr-boilerplate/env"
	"github.com/lokks307/adr-boilerplate/types/e"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	mssql "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mssql/driver"
	mysql "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

var domainConn *sql.DB

func DBLoad() error {
	if err := initializeDB(); err != nil {
		return err
	}

	return nil
}

func initializeDB() error {
	if env.Database.Host == "" || env.Database.User == "" || env.Database.Password == "" || env.Database.DBName == "" || env.Database.Port == 0 {
		return e.PreloadErrInitDBEmptySetting
	}

	dsn := ""

	switch env.Database.Type {
	case "mysql":
		dsn = mysql.MySQLBuildQueryString(env.Database.User, env.Database.Password, env.Database.DBName, env.Database.Host, env.Database.Port, env.Database.SSLMode)
	case "mssql":
		dsn = mssql.MSSQLBuildQueryString(env.Database.User, env.Database.Password, env.Database.DBName, env.Database.Host, env.Database.Port, env.Database.SSLMode)
	case "sqlite3":
		dsn = env.Database.Host
	default:
		return e.PreloadErrInitDBTypeUnsupport
	}

	var dbErr error
	domainConn, dbErr = sql.Open(env.Database.Type, dsn)
	if dbErr != nil {
		logrus.Error(dbErr)
		return e.PreloadErrInitDBConnFailed
	}

	boil.SetDB(domainConn)
	return nil
}

func Conn() *sql.DB {
	return domainConn
}

type InTransaction func(tx *sql.Tx) error

func DoInTransaction(exec *sql.DB, fn InTransaction) error {
	ctx := context.Background()

	tx, err := exec.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err = fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
