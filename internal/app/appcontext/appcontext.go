package appcontext

import (
	"errors"

	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/driver"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
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
	return driver.NewCache(a.getCacheOption())
}

func (a *AppContext) getCacheOption() driver.CacheOption {
	return driver.CacheOption{
		Host:               a.config.GetString("cache.host"),
		Port:               a.config.GetInt("cache.port"),
		Namespace:          a.config.GetString("cache.namespace"),
		Password:           a.config.GetString("cache.password"),
		DialConnectTimeout: a.config.GetDuration("cache.dial_connect_timeout"),
		ReadTimeout:        a.config.GetDuration("cache.read_timeout"),
		WriteTimeout:       a.config.GetDuration("cache.write_timeout"),
		IdleTimeout:        a.config.GetDuration("cache.idle_timeout"),
		MaxConnLifetime:    a.config.GetDuration("cache.conn_lifetime_max"),
		MaxIdle:            a.config.GetInt("cache.conn_idle_max"),
		MaxActive:          a.config.GetInt("cache.conn_active_max"),
		Wait:               a.config.GetBool("cache.is_wait"),
	}
}

// GetInfluxDBClient get Influx DB client
func (a *AppContext) GetInfluxDBClient() (c *influx.Client, err error) {
	influxConfig := influx.ClientConfig{
		Addr:               a.config.GetString("influx.host"),
		Username:           a.config.GetString("influx.user"),
		Password:           a.config.GetString("influx.pass"),
		Database:           a.config.GetString("influx.name"),
		RetentionPolicy:    a.config.GetString("influx.retention_policy"),
		Timeout:            a.config.GetDuration("influx.timeout"),
		InsecureSkipVerify: a.config.GetBool("influx.insecure_skip_verify"),
	}

	return influx.NewClient(influxConfig)
}
