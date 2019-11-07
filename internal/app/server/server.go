package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	plog "github.com/kitabisa/perkakas/v2/log"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/gorp.v2"
)

// IServer interface for server
type IServer interface {
	StartApp()
	StartMetric()
}

type server struct {
	config    config.Provider
	service   *service.Service
	dbMysql   *gorp.DbMap
	dbPostgre *gorp.DbMap
	cachePool *redis.Pool
	logger    *plog.Logger
}

// NewServer create object server
func NewServer(config config.Provider, service *service.Service, dbMysql *gorp.DbMap, dbPostgre *gorp.DbMap, cachePool *redis.Pool, logger *plog.Logger) IServer {
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

		s.logger.AddMessage(plog.InfoLevel, "[API] Server is shutting down")
		s.logger.Print()

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			s.logger.AddMessage(plog.InfoLevel, fmt.Sprintf("[API] Fail to shutting down: %v", err))
			s.logger.Print()
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.config.GetString("app.host"), s.config.GetInt("app.port"))
	srv.Handler = Router(s.service, s.dbMysql, s.dbPostgre, s.cachePool, s.logger)

	s.logger.AddMessage(plog.InfoLevel, fmt.Sprintf("[API] HTTP serve at %s\n", srv.Addr))
	s.logger.Print()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		s.logger.AddMessage(plog.InfoLevel, fmt.Sprintf("[API] Fail to start listen and server: %v", err))
		s.logger.Print()
	}

	<-idleConnectionClosed

	s.logger.AddMessage(plog.InfoLevel, "[API] Bye")
	s.logger.Print()
}

func (s *server) StartMetric() {
	chiMux := chi.NewRouter()
	chiMux.Get("/metrics", prometheus.Handler().ServeHTTP)

	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		s.logger.AddMessage(plog.InfoLevel, "[Metric] Server is shutting down")
		s.logger.Print()

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			s.logger.AddMessage(plog.InfoLevel, fmt.Sprintf("[Metric] Fail to shutting down: %v", err))
			s.logger.Print()
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.config.GetString("app.host"), s.config.GetInt("metric.port"))
	srv.Handler = chiMux

	s.logger.AddMessage(plog.InfoLevel, fmt.Sprintf("[Metric] HTTP serve at %s\n", srv.Addr))
	s.logger.Print()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		s.logger.AddMessage(plog.InfoLevel, fmt.Sprintf("[Metric] Fail to start listen and server: %v", err))
		s.logger.Print()
	}

	<-idleConnectionClosed

	s.logger.AddMessage(plog.InfoLevel, "[Metric] Bye")
	s.logger.Print()
}
