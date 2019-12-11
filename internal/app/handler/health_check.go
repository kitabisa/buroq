package handler

import (
	"net/http"
)

// HealthCheckHandler object for health check handler
type HealthCheckHandler struct {
	HandlerOption
	http.Handler
}

// HealthCheck checking if all work well
func (h HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) (data interface{}, pageToken *string, err error) {
	if h.HandlerOption.Config.GetBool("mysql.is_enabled") {
		err = h.Services.HealthCheck.HealthCheckDbMysql()
		if err != nil {
			return
		}
	}

	if h.HandlerOption.Config.GetBool("postgre.is_enabled") {
		err = h.Services.HealthCheck.HealthCheckDbPostgres()
		if err != nil {
			return
		}
	}

	if h.HandlerOption.Config.GetBool("cache.is_enabled") {
		err = h.Services.HealthCheck.HealthCheckDbCache()
		if err != nil {
			return
		}
	}

	if h.HandlerOption.Config.GetBool("influx.is_enabled") {
		err = h.Services.HealthCheck.HealthCheckInflux()
		if err != nil {
			return
		}
	}

	return
}
