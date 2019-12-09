package server

import (
	"net/http"

	"github.com/go-chi/chi"
	cmiddleware "github.com/go-chi/chi/middleware"
	"github.com/kitabisa/buroq/internal/app/commons"
	"github.com/kitabisa/buroq/internal/app/handler"
	"github.com/kitabisa/buroq/version"
	phttp "github.com/kitabisa/perkakas/v2/http"
	pmiddleware "github.com/kitabisa/perkakas/v2/middleware"
	pstructs "github.com/kitabisa/perkakas/v2/structs"
)

// Router a chi mux
func Router(opt handler.HandlerOption) *chi.Mux {
	handlerCtx := phttp.NewContextHandler(pstructs.Meta{
		Version: version.Version,
		Status:  "stable", //TODO: ask infra if this is used
		APIEnv:  version.Environment,
	})
	commons.InjectErrors(&handlerCtx)

	logMiddleware := pmiddleware.NewHttpRequestLogger(opt.Logger)
	// headerCheckMiddleware := pmiddleware.NewHeaderCheck(handlerCtx, cfg.GetString("app.secret"))

	r := chi.NewRouter()
	// A good base middleware stack (from chi) + middleware from perkakas
	r.Use(cmiddleware.RequestID)
	r.Use(cmiddleware.RealIP)
	r.Use(logMiddleware)
	// r.Use(headerCheckMiddleware) //use this if you want to use default kitabisa's header
	r.Use(cmiddleware.Recoverer)

	// the handler
	phandler := phttp.NewHttpHandler(handlerCtx)

	healthCheckHandler := handler.HealthCheckHandler{}

	healthCheckHandler.HandlerOption = opt
	healthCheckHandler.Handler = phandler(healthCheckHandler.HealthCheck)

	// Setup your routing here
	r.Method(http.MethodGet, "/health_check", healthCheckHandler)
	return r
}

// TODO: func authRouter()
