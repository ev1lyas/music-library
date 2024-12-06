package middleware

import (
	"log"
	_ "music-library/docs" // Документация для сваггера
	"net/http"
	"time"
)

// LoggingMiddleware - Логирует все входящие запросы
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Запрос: %s %s, IP: %s, Время: %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware - Добавляет заголовки CORS
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
