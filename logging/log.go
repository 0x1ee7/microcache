package logging

import (
	"log"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

// Handler wraps other handler to enable logging.
func Handler(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rec := &statusRecorder{
				ResponseWriter: w,
				Status:         200,
			}
			next.ServeHTTP(rec, r)
			logger.Println(r.Method, r.URL.Path, r.RemoteAddr, rec.Status, r.UserAgent())
		})
	}
}
