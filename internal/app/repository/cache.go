package repository

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/im7mortal/kmutex"
)

// ICacheRepo interface for cache repo
type ICacheRepo interface {
	WriteCache(key string, data interface{}, ttl time.Duration) (err error)
	WriteCacheIfEmpty(key string, data interface{}, ttl time.Duration) (err error)
}

type cacheRepo struct {
	opt    Option
	kmutex *kmutex.Kmutex
}

// NewCacheRepository initiate cache repo
func NewCacheRepository(opt Option) ICacheRepo {
	return &cacheRepo{
		opt:    opt,
		kmutex: kmutex.New(),
	}
}

// WriteCache this will and must write the data to cache with corresponding key using locking
func (c *cacheRepo) WriteCache(key string, data interface{}, ttl time.Duration) (err error) {
	c.kmutex.Lock(key)
	defer c.kmutex.Unlock(key)

	// write data to cache
	conn := c.opt.CachePool.Get()
	_, err = conn.Do("SETEX", key, ttl.Seconds(), data)
	if err != nil {
		return err
	}

	return nil
}

// WriteCacheIfEmpty will try to write to cache, if the data still empty after locking
func (c *cacheRepo) WriteCacheIfEmpty(key string, data interface{}, ttl time.Duration) (err error) {
	c.kmutex.Lock(key)
	defer c.kmutex.Unlock(key)

	// check whether cache value is empty
	conn := c.opt.CachePool.Get()
	_, err = conn.Do("GET", key)
	if err != nil {
		if err == redis.ErrNil {
			return nil //return nil as the data already set, no need to overwrite
		}

		return err
	}

	// write data to cache
	_, err = conn.Do("SETEX", key, ttl.Seconds(), data)
	if err != nil {
		return err
	}

	return nil
}
