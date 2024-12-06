package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB — глобальная переменная для работы с базой данных
var DB *gorm.DB

// InitDB инициализирует подключение к базе данных
func InitDB() {
	// Получение параметров подключения из переменных окружения
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5432")
	user := getEnvOrDefault("DB_USER", "postgres")
	password := getEnvOrDefault("DB_PASSWORD", "")
	dbname := getEnvOrDefault("DB_NAME", "testdb")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}
	log.Println("Подключение к базе данных успешно")

	// Создание и заполнение таблицы songs
	if err := setupDatabase(DB); err != nil {
		log.Fatalf("Ошибка настройки базы данных: %v", err)
	}
}

// getEnvOrDefault возвращает значение переменной окружения или значение по умолчанию
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// setupDatabase выполняет миграции и заполняет тестовыми данными
func setupDatabase(db *gorm.DB) error {
	log.Println("Выполнение миграций...")
	if err := ExecSQLFile(db, "migrations/001_create_songs.sql"); err != nil {
		return fmt.Errorf("ошибка при создании таблиц: %w", err)
	}

	/* log.Println("Заполнение тестовыми данными...")
	if err := ExecSQLFile(db, "migrations/002_insert_songs.sql"); err != nil {
		return fmt.Errorf("ошибка при заполнении тестовыми данными: %w", err)
	} */

	return nil
}

// CleanUpTestData удаляет тестовые данные после завершения программы
func CleanUpTestData() {
	if err := DB.Exec("DELETE FROM songs").Error; err != nil {
		log.Printf("Ошибка при удалении тестовых данных: %v", err)
	} else {
		log.Println("Тестовые данные успешно удалены.")
	}
}
