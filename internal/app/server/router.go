package server

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
)

func Router(service service.Service) *chi.Mux {
	r := chi.NewRouter()

	// TODO: setup middleware
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)

	// TODO: use our perkakas logger instead
	r.Use(chiMiddleware.Logger)

	r.Use(chiMiddleware.Recoverer)

	// Setup your routing here
	// r.Get("/health_check")
	return r
}
