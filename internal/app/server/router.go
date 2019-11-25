package server

import (
	"github.com/go-chi/chi"
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/handler"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/go-bootstrap/version"
	"github.com/kitabisa/perkakas/v2/distlock"
	phttp "github.com/kitabisa/perkakas/v2/http"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/middleware"
	pstructs "github.com/kitabisa/perkakas/v2/structs"
	"gopkg.in/gorp.v2"
)

// Router a chi mux
func Router(cfg config.Provider, service *service.Service, dbMysql *gorp.DbMap, dbPostgre *gorp.DbMap, cachePool *redis.Pool, cacheDistLock *distlock.DistLock, logger *log.Logger) *chi.Mux {
	logMiddleware := middleware.NewHttpRequestLogger(logger)
	handlerCtx := phttp.NewContextHandler(pstructs.Meta{
		Version: "123",
		Status:  "stable", //TODO: ask infra if this is used
		APIEnv:  version.Environment,
	})
	headerCheckMiddleware := middleware.NewHeaderCheck(handlerCtx, cfg.GetString("app.secret"))
	r := chi.NewRouter()

	// TODO: use perkakas handler

	// use middlewares needed
	// r.Use(chiMiddleware.RequestID)
	// r.Use(chiMiddleware.RealIP)
	r.Use(logMiddleware)
	r.Use(headerCheckMiddleware)
	// r.Use(chiMiddleware.Recoverer)

	// Setup your routing here
	// r.Get("/health_check")
	h := handler.NewHandler(service, dbMysql, dbPostgre, cachePool, cacheDistLock, logger)
	r.Get("/health_check", h.HealthCheck())
	return r
}

// TODO: func authRouter()
