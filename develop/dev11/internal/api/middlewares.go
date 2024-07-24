package api

import (
	"log"
	"net/http"
)

func LoggerMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("New request from %s\n{\n\tMethod: %s | API-path: %s | User-Agent: %s\n}\n\n",
			r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
		handler(w, r)
	}
}
