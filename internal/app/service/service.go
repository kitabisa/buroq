package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/internal/app/repository"
	"github.com/kitabisa/perkakas/v2/distlock"
	"gopkg.in/gorp.v2"
)

// Option anything any service object needed
type Option struct {
	DbMysql       *gorp.DbMap
	DbPostgre     *gorp.DbMap
	CachePool     *redis.Pool
	CacheDistLock *distlock.DistLock
	Repo          *repository.Repository
}

// Service all service object injected here
type Service struct {
	HealthCheck IHealthCheck
}

// NewService preparing empty object of Service
func NewService() *Service {
	return &Service{}
}

// TODO: set function for each service
// eg
/*
func (s *Service) SetUserService(userService IUserService) {
	s.User = userService
}
*/
