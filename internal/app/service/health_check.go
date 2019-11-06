package service

import (
	"fmt"

	"github.com/kitabisa/go-bootstrap/internal/pkg/commons"
)

type IHealthCheck interface {
	HealthCheck() (err error, rc int)
}

type healthCheck struct {
	option Option
}

func NewHealthCheck(option Option) IHealthCheck {
	return &healthCheck{
		option: option,
	}
}

func (h *healthCheck) HealthCheck() (err error, rc int) {
	err = h.option.DbMysql.Db.Ping()
	if err != nil {
		// TODO: logging
		fmt.Println(err)
		rc = commons.RCDBConnectionError
		return
	}

	err = h.option.DbPostgre.Db.Ping()
	if err != nil {
		// TODO: logging
		fmt.Println(err)
		rc = commons.RCDBConnectionError
		return
	}

	cacheConn := h.option.CachePool.Get()
	_, err = cacheConn.Do("PING")
	if err != nil {
		// TODO: logging
		rc = commons.RCCacheConnectionError
		return
	}
	defer cacheConn.Close()

	return nil, commons.RCSuccess
}
