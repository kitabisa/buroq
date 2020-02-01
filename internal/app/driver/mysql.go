package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // defines mysql driver used
	"gopkg.in/gorp.v2"
)

// DBMysqlOption options for mysql connection
type DBMysqlOption struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

// NewMysqlDatabase return gorp dbmap object with MySQL options param
func NewMysqlDatabase(option DBMysqlOption) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.Username, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(option.ConnMaxLifetime)
	db.SetMaxIdleConns(option.MaxIdleConns)
	db.SetMaxOpenConns(option.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	gorp := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return gorp, nil
}
