package driver

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// CacheOption properties for cache DB
type CacheOption struct {
	Host               string
	Port               int
	DialConnectTimeout time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxIdle            int
	MaxActive          int
	IdleTimeout        time.Duration
	Wait               bool
	MaxConnLifetime    time.Duration
	Password           string
	Namespace          string
}

// NewCache create cache pool
func NewCache(option CacheOption) *redis.Pool {
	dialConnectTimeoutOption := redis.DialConnectTimeout(option.DialConnectTimeout)
	readTimeoutOption := redis.DialReadTimeout(option.ReadTimeout)
	writeTimeoutOption := redis.DialWriteTimeout(option.WriteTimeout)

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(fmt.Sprintf("redis://%s@%s:%d", option.Password, option.Host, option.Port), dialConnectTimeoutOption, readTimeoutOption, writeTimeoutOption)
			if err != nil {
				return nil, fmt.Errorf("ERROR connect redis | %v", err)
			}

			if option.Password != "" {
				if _, err := c.Do("AUTH", option.Password); err != nil {
					return nil, fmt.Errorf("ERROR on AUTH redis | %v", err)
				}
			}

			if _, err := c.Do("SELECT", option.Namespace); err != nil {
				return nil, fmt.Errorf("ERROR on SELECT namespace redis | %v", err)
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return nil
		},
		MaxIdle:         option.MaxIdle,
		MaxActive:       option.MaxActive,
		IdleTimeout:     option.IdleTimeout,
		Wait:            option.Wait,
		MaxConnLifetime: option.MaxConnLifetime,
	}

	return pool
}
