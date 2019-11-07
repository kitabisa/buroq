package middlewares

import (
	plog "github.com/kitabisa/perkakas/v2/log"
)

// Middleware object
type Middleware struct {
	logger *plog.Logger
}

func NewMiddleware(logger *plog.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}
