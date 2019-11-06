package server

import (
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/kitabisa/go-bootstrap/internal/app/handler"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/go-bootstrap/internal/pkg/appcontext"
)

func Router(appCtx *appcontext.AppContext, service *service.Service) *chi.Mux {
	r := chi.NewRouter()

	// TODO: setup middleware
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)

	// TODO: use our perkakas logger instead
	r.Use(chiMiddleware.Logger)

	r.Use(chiMiddleware.Recoverer)

	// Setup your routing here
	// r.Get("/health_check")
	h := handler.NewHandler(service, appCtx)
	r.MethodFunc(http.MethodGet, "/health_check", h.HealthCheck)
	return r
}
