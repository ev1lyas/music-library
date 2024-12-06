package main

import (
	"io"
	"log"
	"music-library/internal/db"
	"music-library/internal/router"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv загружает и проверяет переменные окружения
func LoadEnv() {
	required := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "SERVER_PORT"}
	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Не найдена обязательная переменная окружения: %s", key)
		}
	}
}

// @title Music Library API
// @version 1.0
// @description API для управления музыкальной библиотекой
// @host localhost:8000
// @BasePath /
func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	// Проверка обязательных переменных окружения
	LoadEnv()

	// Логирование в файл и консоль
	file, err := os.OpenFile("debug/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл для логов: %v", err)
	}
	defer file.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	// Инициализация базы данных
	db.InitDB()

	// Настройка маршрутов
	r := router.SetupRouter()

	// Чтение порта из окружения
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000" // Значение по умолчанию
	}

	// Запуск сервера
	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
