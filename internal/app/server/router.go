package server

import (
	"github.com/go-chi/chi"
	cmiddleware "github.com/go-chi/chi/middleware"
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/handler"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/go-bootstrap/internal/pkg/commons"
	"github.com/kitabisa/go-bootstrap/version"
	phttp "github.com/kitabisa/perkakas/v2/http"
	"github.com/kitabisa/perkakas/v2/log"
	pmiddleware "github.com/kitabisa/perkakas/v2/middleware"
	pstructs "github.com/kitabisa/perkakas/v2/structs"
	"gopkg.in/gorp.v2"
	"net/http"
)

// Router a chi mux
func Router(cfg config.Provider, service *service.Service, dbMysql *gorp.DbMap, dbPostgre *gorp.DbMap, cachePool *redis.Pool, logger *log.Logger) *chi.Mux {
	handlerCtx := phttp.NewContextHandler(pstructs.Meta{
		Version: version.Version,
		Status:  "stable", //TODO: ask infra if this is used
		APIEnv:  version.Environment,
	})
	commons.InjectErrors(&handlerCtx)

	logMiddleware := pmiddleware.NewHttpRequestLogger(logger)
	// headerCheckMiddleware := pmiddleware.NewHeaderCheck(handlerCtx, cfg.GetString("app.secret"))

	r := chi.NewRouter()
	// A good base middleware stack (from chi) + middleware from perkakas
	r.Use(cmiddleware.RequestID)
	r.Use(cmiddleware.RealIP)
	r.Use(logMiddleware)
	// r.Use(headerCheckMiddleware)
	r.Use(cmiddleware.Recoverer)

	// the handler
	phandler := phttp.NewHttpHandler(handlerCtx)
	handlerOpt := handler.HandlerOption{
		Config:    cfg,
		Services:  service,
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cachePool,
		Logger:    logger,
	}

	healthCheckHandler := handler.HealthCheckHandler{}

	healthCheckHandler.HandlerOption = handlerOpt
	healthCheckHandler.Handler = phandler(healthCheckHandler.HealthCheck)
	// ph := phttp.NewHttpHandler(handlerCtx)

	// Setup your routing here
	r.Method(http.MethodGet, "/health_check", healthCheckHandler)
	return r
}

// TODO: func authRouter()
