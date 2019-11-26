package handler

import (
	"net/http"
)

type HealthCheckHandler struct {
	HandlerOption
	http.Handler
}

// HealthCheck checking if all work well
func (h HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) (data interface{}, pageToken *string, err error) {
	err = h.Services.HealthCheck.HealthCheckDbMysql()
	if err != nil {
		return
	}

	err = h.Services.HealthCheck.HealthCheckDbPostgres()
	if err != nil {
		return
	}

	err = h.Services.HealthCheck.HealthCheckDbCache()
	if err != nil {
		return
	}

	return
}
