package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"music-library/internal/models"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// ErrorResponse представляет структуру ошибки
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageResponse представляет ответ с сообщением
type MessageResponse struct {
	Message string `json:"message"`
}

// HandleError - универсальная обработка ошибок
func HandleError(w http.ResponseWriter, message string, statusCode int, err error) {
	if err != nil {
		log.Printf("Ошибка: %s, Подробности: %v\n", message, err)
	}
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// ParsePaginationParams - Парсинг параметров пагинации
func ParsePaginationParams(pageStr, limitStr string, defaultPage, defaultLimit int) (page, limit int) {
	page = defaultPage
	limit = defaultLimit

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	return
}

// PaginateSlice - Пагинация для срезов
func PaginateSlice(data []string, page, limit int) []string {
	start := (page - 1) * limit
	if start >= len(data) {
		return nil
	}

	end := start + limit
	if end > len(data) {
		end = len(data)
	}

	return data[start:end]
}

// ApplyUpdates - Применение обновлений к структуре
func ApplyUpdates(song *models.Song, updates models.UpdateSongRequest) {
	if updates.MusicGroup != nil {
		song.MusicGroup = *updates.MusicGroup
	}
	if updates.Song != nil {
		song.Song = *updates.Song
	}
	if updates.ReleaseDate != nil {
		song.ReleaseDate = *updates.ReleaseDate
	}
	if updates.Text != nil {
		song.Text = *updates.Text
	}
	if updates.Link != nil {
		song.Link = *updates.Link
	}
}

// FetchSongDetails - Получение данных песни из внешнего API
func FetchSongDetails(group, song string) (*models.SongDetail, error) {
	// Устанавливаем таймаут для HTTP-запроса
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Получаем URL внешнего API из переменных окружения
	baseURL := os.Getenv("EXTERNAL_API_URL")
	if baseURL == "" {
		return nil, errors.New("переменная окружения EXTERNAL_API_URL не установлена")
	}

	// Формируем URL для API
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга базового URL: %w", err)
	}
	query := apiURL.Query()
	query.Set("group", group)
	query.Set("song", song)
	apiURL.RawQuery = query.Encode()

	// Логируем конечный URL запроса
	log.Printf("Запрос к внешнему API: %s", apiURL.String())

	// Создаем HTTP-запрос
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания HTTP-запроса: %w", err)
	}

	// Выполняем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения HTTP-запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка API: статус %d", resp.StatusCode)
	}

	// Парсим ответ
	var result struct {
		ReleaseDate string `json:"releaseDate"`
		Text        string `json:"text"`
		Link        string `json:"link"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Проверяем наличие обязательных данных
	if result.ReleaseDate == "" || result.Text == "" || result.Link == "" {
		return nil, errors.New("ответ API не содержит необходимых данных")
	}

	// Возвращаем структуру SongDetail
	return &models.SongDetail{
		ReleaseDate: result.ReleaseDate,
		Text:        result.Text,
		Link:        result.Link,
	}, nil
}
