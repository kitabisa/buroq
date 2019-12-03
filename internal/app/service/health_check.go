package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/internal/app/commons"
	"github.com/kitabisa/go-bootstrap/internal/app/repository"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
	"gopkg.in/gorp.v2"
)

// IHealthCheck interface for health check service
type IHealthCheck interface {
	HealthCheckDbMysql() (err error)
	HealthCheckDbPostgres() (err error)
	HealthCheckDbCache() (err error)
	HealthCheckInflux() (err error)
}

type healthCheck struct {
	Repo      *repository.Repository
	dbMysql   *gorp.DbMap
	dbPostgre *gorp.DbMap
	cachePool *redis.Pool
	influx    *influx.Client
	logger    *log.Logger
}

// NewHealthCheck create health check service instance with option as param
func NewHealthCheck(option Option) IHealthCheck {
	return &healthCheck{
		Repo:      option.Repo,
		dbMysql:   option.DbMysql,
		dbPostgre: option.DbPostgre,
		cachePool: option.CachePool,
		influx:    option.Influx,
		logger:    option.Logger,
	}
}

func (h *healthCheck) HealthCheckDbMysql() (err error) {
	err = h.dbMysql.Db.Ping()
	if err != nil {
		h.logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbPostgres() (err error) {
	err = h.dbPostgre.Db.Ping()
	if err != nil {
		h.logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbCache() (err error) {
	cacheConn := h.cachePool.Get()
	_, err = cacheConn.Do("PING")
	if err != nil {
		h.logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrCacheConn
		return
	}
	defer cacheConn.Close()

	return nil
}

func (h *healthCheck) HealthCheckInflux() (err error) {
	err = h.influx.Ping()
	if err != nil {
		h.logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrInfluxConn
	}

	return
}
