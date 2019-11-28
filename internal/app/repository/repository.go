package repository

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
	"gopkg.in/gorp.v2"
)

// Option anything any repo object needed
type Option struct {
	DbMysql   *gorp.DbMap
	DbPostgre *gorp.DbMap
	CachePool *redis.Pool
	Influx    *influx.Client
	Logger    *log.Logger
}

// Repository all repo object injected here
type Repository struct {
	// User IUserRepository
	Cache ICacheRepo
}
