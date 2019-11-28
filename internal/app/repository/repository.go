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
	cache ICacheRepo
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

// SetCacheRepo Set Cache Repository to Repository collection
func (r *Repository) SetCacheRepo(cacheRepo ICacheRepo) {
	r.cache = cacheRepo
}
