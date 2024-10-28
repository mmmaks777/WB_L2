package server

import (
	"log"
	"net/http"
	"time"
)

func LoggingMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Start request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("End request: %s %s, duration: %s", r.Method, r.URL.Path, time.Since(startTime))
	})
}
