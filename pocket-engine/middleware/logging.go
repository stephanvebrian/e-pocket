package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func StartLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming Request Start", r.RequestURI)

		ctx := context.WithValue(r.Context(), startTimeKey, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func EndLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler
		next.ServeHTTP(w, r)

		// Retrieve the start time from the context
		startTime, ok := r.Context().Value(startTimeKey).(time.Time)
		duration := "-"
		if ok {
			// Calculate the duration if the start time exists
			duration = fmt.Sprintf("%d ms", time.Since(startTime).Milliseconds())
		}

		// Log the request URI and the duration in a single line
		log.Printf("Request: %s, Duration: %s", r.RequestURI, duration)
	})
}
