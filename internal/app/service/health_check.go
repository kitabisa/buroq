package service

import (
	"github.com/kitabisa/buroq/internal/app/commons"
	"github.com/kitabisa/perkakas/v2/log"
)

// IHealthCheck interface for health check service
type IHealthCheck interface {
	HealthCheckDbMysql() (err error)
	HealthCheckDbPostgres() (err error)
	HealthCheckDbCache() (err error)
	HealthCheckInflux() (err error)
}

type healthCheck struct {
	opt Option
}

// NewHealthCheck create health check service instance with option as param
func NewHealthCheck(opt Option) IHealthCheck {
	return &healthCheck{
		opt: opt,
	}
}

func (h *healthCheck) HealthCheckDbMysql() (err error) {
	err = h.opt.DbMysql.Db.Ping()
	if err != nil {
		h.opt.Logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbPostgres() (err error) {
	err = h.opt.DbPostgre.Db.Ping()
	if err != nil {
		h.opt.Logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbCache() (err error) {
	cacheConn := h.opt.CachePool.Get()
	_, err = cacheConn.Do("PING")
	if err != nil {
		h.opt.Logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrCacheConn
		return
	}
	defer cacheConn.Close()

	return nil
}

func (h *healthCheck) HealthCheckInflux() (err error) {
	err = h.opt.Influx.Ping()
	if err != nil {
		h.opt.Logger.AddMessage(log.FatalLevel, err.Error())
		err = commons.ErrInfluxConn
	}

	return
}
