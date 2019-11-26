package service

import (
	"github.com/kitabisa/go-bootstrap/internal/pkg/commons"
	"github.com/kitabisa/perkakas/v2/log"
)

// IHealthCheck interface for health check service
type IHealthCheck interface {
	HealthCheckDbMysql() (err error)
	HealthCheckDbPostgres() (err error)
	HealthCheckDbCache() (err error)
}

type healthCheck struct {
	Option
}

// NewHealthCheck create health check service instance with option as param
func NewHealthCheck(option Option) IHealthCheck {
	return &healthCheck{
		option,
	}
}

func (h *healthCheck) HealthCheckDbMysql() (err error) {
	err = h.DbMysql.Db.Ping()
	if err != nil {
		h.Logger.AddMessage(log.FatalLevel, err.Error()).Print()
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbPostgres() (err error) {
	err = h.DbPostgre.Db.Ping()
	if err != nil {
		h.Logger.AddMessage(log.FatalLevel, err.Error()).Print()
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbCache() (err error) {
	cacheConn := h.CachePool.Get()
	_, err = cacheConn.Do("PING")
	if err != nil {
		h.Logger.AddMessage(log.FatalLevel, err.Error()).Print()
		err = commons.ErrCacheConn
		return
	}
	defer cacheConn.Close()

	return nil
}
