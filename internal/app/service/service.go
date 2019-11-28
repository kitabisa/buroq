package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/internal/app/repository"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
	"gopkg.in/gorp.v2"
)

// Option anything any service object needed
type Option struct {
	DbMysql   *gorp.DbMap
	DbPostgre *gorp.DbMap
	CachePool *redis.Pool
	Repo      *repository.Repository
	Influx    *influx.Client
	Logger    *log.Logger
}

// Service all service object injected here
type Service struct {
	HealthCheck IHealthCheck
}

// NewService preparing empty object of Service
func NewService() *Service {
	return &Service{}
}

// SetHealthCheckService set health check service
func (s *Service) SetHealthCheckService(hc IHealthCheck) {
	s.HealthCheck = hc
}
