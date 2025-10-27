package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		w.Header().Get("Status Code")
		log.Println(w.Header().Get("Status Code"), r.Method, r.URL.Path)
	})

}
