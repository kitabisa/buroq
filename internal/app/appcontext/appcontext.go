package appcontext

import (
	"errors"

	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/driver"
	"gopkg.in/gorp.v2"
)

const (
	// DBDialectMysql rdbms dialect name for MySQL
	DBDialectMysql = "mysql"

	// DBDialectPostgres rdbms dialect name for PostgreSQL
	DBDialectPostgres = "postgres"
)

// AppContext the app context struct
type AppContext struct {
	config config.Provider
}

// NewAppContext initiate appcontext object
func NewAppContext(config config.Provider) *AppContext {
	return &AppContext{
		config: config,
	}
}

// GetDBInstance getting gorp instance, param: dbType can be "mysql" or "postgre"
func (a *AppContext) GetDBInstance(dbType string) (*gorp.DbMap, error) {
	var gorp *gorp.DbMap
	var err error
	switch dbType {
	case DBDialectMysql:
		dbOption := a.getMysqlOption()
		gorp, err = driver.NewMysqlDatabase(dbOption)
	case DBDialectPostgres:
		dbOption := a.getPostgreOption()
		gorp, err = driver.NewPostgreDatabase(dbOption)
	default:
		err = errors.New("Error get db instance, unknown db type")
	}

	return gorp, err
}

func (a *AppContext) getMysqlOption() driver.DBMysqlOption {
	return driver.DBMysqlOption{
		Host:                 a.config.GetString("mysql.host"),
		Port:                 a.config.GetInt("mysql.port"),
		Username:             a.config.GetString("mysql.username"),
		Password:             a.config.GetString("mysql.password"),
		DBName:               a.config.GetString("mysql.name"),
		AdditionalParameters: a.config.GetString("mysql.additional_parameters"),
		MaxOpenConns:         a.config.GetInt("mysql.conn_open_max"),
		MaxIdleConns:         a.config.GetInt("mysql.conn_idle_max"),
		ConnMaxLifetime:      a.config.GetDuration("mysql.conn_lifetime_max"),
	}
}

func (a *AppContext) getPostgreOption() driver.DBPostgreOption {
	return driver.DBPostgreOption{
		Host:        a.config.GetString("postgre.host"),
		Port:        a.config.GetInt("postgre.port"),
		Username:    a.config.GetString("postgre.username"),
		Password:    a.config.GetString("postgre.password"),
		DBName:      a.config.GetString("postgre.name"),
		MaxPoolSize: a.config.GetInt("postgre.pool_size"),
	}
}

// GetCachePool get cache pool connection
func (a *AppContext) GetCachePool() *redis.Pool {
	return driver.NewRedis(a.getCacheOption())
}

func (a *AppContext) getCacheOption() driver.RedisOption {
	return driver.RedisOption{
		Host:               a.config.GetString("redis.host"),
		Port:               a.config.GetInt("redis.port"),
		Namespace:          a.config.GetString("redis.namespace"),
		Password:           a.config.GetString("redis.password"),
		DialConnectTimeout: a.config.GetDuration("redis.dial_connect_timeout"),
		ReadTimeout:        a.config.GetDuration("redis.read_timeout"),
		WriteTimeout:       a.config.GetDuration("redis.write_timeout"),
		IdleTimeout:        a.config.GetDuration("redis.idle_timeout"),
		MaxConnLifetime:    a.config.GetDuration("redis.conn_lifetime_max"),
		MaxIdle:            a.config.GetInt("redis.conn_idle_max"),
		MaxActive:          a.config.GetInt("redis.conn_active_max"),
		Wait:               a.config.GetBool("redis.is_wait"),
	}
}
