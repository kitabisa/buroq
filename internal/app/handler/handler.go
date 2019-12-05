package handler

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/buroq/config"
	"github.com/kitabisa/buroq/internal/app/service"
	"github.com/kitabisa/perkakas/v2/log"
	"gopkg.in/gorp.v2"
)

type HandlerOption struct {
	Config    config.Provider
	Services  *service.Service
	DbMysql   *gorp.DbMap
	DbPostgre *gorp.DbMap
	CachePool *redis.Pool
	Logger    *log.Logger
}
