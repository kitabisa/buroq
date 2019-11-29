package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorp.v2"
)

// IServer interface for server
type IServer interface {
	StartApp()
}

type server struct {
	config    config.Provider
	service   *service.Service
	dbMysql   *gorp.DbMap
	dbPostgre *gorp.DbMap
	cachePool *redis.Pool
	logger    *log.Logger
}

// NewServer create object server
func NewServer(config config.Provider, service *service.Service, dbMysql *gorp.DbMap, dbPostgre *gorp.DbMap, cachePool *redis.Pool, logger *log.Logger) IServer {
	return &server{
		config:    config,
		service:   service,
		dbMysql:   dbMysql,
		dbPostgre: dbPostgre,
		cachePool: cachePool,
		logger:    logger,
	}
}

func (s *server) StartApp() {
	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		logrus.Infoln("[API] Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logrus.Infof("[API] Fail to shutting down: %v", err)
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.config.GetString("app.host"), s.config.GetInt("app.port"))
	srv.Handler = Router(s.config, s.service, s.dbMysql, s.dbPostgre, s.cachePool, s.logger)

	logrus.Infof("[API] HTTP serve at %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logrus.Infof("[API] Fail to start listen and server: %v", err)
	}

	<-idleConnectionClosed
	logrus.Infoln("[API] Bye")
}
