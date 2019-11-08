package repository

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/perkakas/v2/distlock"
	"github.com/kitabisa/perkakas/v2/log"
	"gopkg.in/gorp.v2"
)

// Option anything any repo object needed
type Option struct {
	DbMysql       *gorp.DbMap
	DbPostgre     *gorp.DbMap
	CachePool     *redis.Pool
	CacheDistLock *distlock.DistLock
	Logger        *log.Logger
}

// Repository all repo object injected here
type Repository struct {
	// User IUserRepository
}

// NewRepository preparing empty object of Repo
func NewRepository() *Repository {
	return &Repository{}
}

// TODO: set function for each repository
// eg
/*
func (r *Repository) SetUserRepository(userRepository IUserRepository) {
	r.User = userRepository
}
*/
