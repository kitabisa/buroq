package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // defines postgreSQL driver used
	"gopkg.in/gorp.v2"
)

// DBPostgreOption options for postgre connection
type DBPostgreOption struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	MaxPoolSize int
}

// NewPostgreDatabase return gorp dbmap object with postgre options param
func NewPostgreDatabase(option DBPostgreOption) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", option.Host, option.Port, option.Username, option.DBName, option.Password))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(option.MaxPoolSize)
	gorp := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return gorp, nil
}
