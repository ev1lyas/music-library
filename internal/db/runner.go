package db

import (
	"bufio"
	"errors"
	"log"
	"os"

	"gorm.io/gorm"
)

// ExecSQLFile читает SQL-файл и выполняет запросы
func ExecSQLFile(db *gorm.DB, filepath string) error {
	// Проверка существования файла
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return errors.New("файл не найден: " + filepath)
	}

	// Открытие файла
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Printf("Чтение SQL-файла: %s", filepath)

	// Использование транзакции
	return db.Transaction(func(tx *gorm.DB) error {
		scanner := bufio.NewScanner(file)
		var query string

		// Чтение файла построчно
		for scanner.Scan() {
			line := scanner.Text()
			query += line + " "

			// Если в строке есть конец SQL-запроса (';'), выполнить запрос
			if len(line) > 0 && line[len(line)-1] == ';' {
				if err := tx.Exec(query).Error; err != nil {
					return err
				}
				query = "" // Очистить запрос после выполнения
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		log.Printf("SQL запросы из файла %s успешно выполнены", filepath)
		return nil
	})
}
