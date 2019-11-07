package middlewares

import (
	"fmt"
	"net/http"

	httputil "github.com/kitabisa/perkakas/v2/httputil"
)

// these header values won't be printed to log
var secretHeaderKeys = []string{"Authorization", "X-Ktbs-Signature"}

// LogRequest middleware for logging request header & body
func (m *Middleware) LogRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var requestStr string
		for k, v := range r.Header {
			requestStr = fmt.Sprintf("%s%s", requestStr, fmt.Sprintf("%s:%s\n", k, v))
		}
		requestStr = fmt.Sprintf("%s%s", requestStr, httputil.ReadRequestBody(r))
		m.logger.SetRequest(requestStr)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
