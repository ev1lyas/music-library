package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"music-library/internal/controllers"
	"music-library/internal/middleware"
)

// SetupRouter настраивает маршруты
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Middleware для логирования и CORS
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)

	// Маршруты для работы с песнями
	r.HandleFunc("/songs", controllers.GetSongs).Methods("GET")
	r.HandleFunc("/songs/{id}/verse", controllers.GetSongVerse).Methods("GET")
	r.HandleFunc("/songs/{id}", controllers.DeleteSong).Methods("DELETE")
	r.HandleFunc("/songs/{id}", controllers.UpdateSong).Methods("PUT")
	r.HandleFunc("/songs", controllers.AddSong).Methods("POST")

	// Swagger-документация
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
