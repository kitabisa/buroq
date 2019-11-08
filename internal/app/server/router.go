package server

import (
	"github.com/go-chi/chi"
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/internal/app/handler"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/middleware"
	"gopkg.in/gorp.v2"
)

// Router a chi mux
func Router(service *service.Service, dbMysql *gorp.DbMap, dbPostgre *gorp.DbMap, cachePool *redis.Pool, logger *log.Logger) *chi.Mux {
	logMiddleware := middleware.NewHttpRequestLogger(logger)
	r := chi.NewRouter()

	// use middlewares needed
	// r.Use(chiMiddleware.RequestID)
	// r.Use(chiMiddleware.RealIP)
	r.Use(logMiddleware)
	// r.Use(chiMiddleware.Recoverer)

	// Setup your routing here
	// r.Get("/health_check")
	h := handler.NewHandler(service, dbMysql, dbPostgre, cachePool, logger)
	r.Get("/health_check", h.HealthCheck())
	return r
}
