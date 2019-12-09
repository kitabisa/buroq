package commons

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/buroq/config"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
	"gopkg.in/gorp.v2"
)

type Options struct {
	Config    config.Provider
	DbMysql   *gorp.DbMap
	DbPostgre *gorp.DbMap
	CachePool *redis.Pool
	Influx    *influx.Client
	Logger    *log.Logger
}
