package middleware

import (
	"net/http"
	"sabigo/config"
)

func AccessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cors_allow := config.LoadEnvironmentalVariables("CROSS_ORIGIN_ALLOW")
		origin := config.LoadEnvironmentalVariables("ALLOWED_ORIGIN")
		if cors_allow == "true" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

			if r.Method == "OPTIONS" {
				return
			}

			next.ServeHTTP(w, r)
		}
	})
}
