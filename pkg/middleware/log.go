package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LogMiddleware(next http.Handler) http.Handler {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &WrapperWriter{
			ResponseWriter: w,
		}
		next.ServeHTTP(wrapper, r)
		logrus.WithFields(logrus.Fields{
			"status": wrapper.StatusCode,
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("Request processed")

	})
}
