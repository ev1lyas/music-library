package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"music-library/internal/db"
	"music-library/internal/models"
	"music-library/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

// GetSongs godoc
// @Summary Получить список песен с фильтрацией и пагинацией
// @Description Возвращает список песен с возможностью фильтрации по полям и пагинации. Для управления фильтрацией и пагинацией необходимо в теле запроса указать limit элементов на странице и page номер страницы.
// @Tags songs
// @Accept json
// @Produce json
// @Param musicGroup query string false "Название музыкальной группы"
// @Param song query string false "Название песни"
// @Param releaseDate query string false "Дата релиза песни"
// @Param text query string false "Текст песни"
// @Param link query string false "Ссылка на источник"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество элементов на странице" default(10)
// @Success 200 {array} models.Song "Список песен"
// @Failure 400 {object} models.ErrorResponse "Неверные параметры запроса"
// @Failure 404 {object} models.ErrorResponse "Песня или страница не найдена"
// @Failure 500 {object} models.ErrorResponse "Ошибка при обработке запроса"
// @Router /songs [get]
func GetSongs(w http.ResponseWriter, r *http.Request) {
	var songs []models.Song
	query := db.DB

	params := r.URL.Query()

	// Фильтрация по параметрам
	filterableFields := map[string]string{
		"musicGroup":  "musicGroup LIKE ?",
		"song":        "song LIKE ?",
		"releaseDate": "releaseDate LIKE ?",
		"text":        "text LIKE ?",
		"link":        "link LIKE ?",
	}

	for param, queryStr := range filterableFields {
		if value := params.Get(param); value != "" {
			query = query.Where(queryStr, "%"+value+"%")
		}
	}

	// Параметры пагинации
	page, limit := utils.ParsePaginationParams(params.Get("page"), params.Get("limit"), 1, 10)
	offset := (page - 1) * limit

	query = query.Offset(offset).Limit(limit)

	if err := query.Find(&songs).Error; err != nil {
		log.Printf("Ошибка при получении песен: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(songs) == 0 {
		log.Printf("Песни не найдены")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Песни не найдены"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

// GetSongVerse godoc
// @Summary Получить текст песни с пагинацией по куплетам
// @Description Возвращает текст песни, разделенный на куплеты, с поддержкой пагинации
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество куплетов на странице" default(1)
// @Success 200 {array} string "Текст песни в виде массива строк"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 404 {object} models.ErrorResponse "Песня не найдена"
// @Failure 500 {object} models.ErrorResponse "Ошибка при обработке запроса"
// @Router /songs/{id}/verse [get]
func GetSongVerse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.HandleError(w, "Неверный ID песни", http.StatusBadRequest, err)
		return
	}

	page, limit := utils.ParsePaginationParams(r.URL.Query().Get("page"), r.URL.Query().Get("limit"), 1, 1)

	var song models.Song
	if err := db.DB.First(&song, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.HandleError(w, "Песня не найдена", http.StatusNotFound, err)
			return
		}
		utils.HandleError(w, "Ошибка при получении песни", http.StatusInternalServerError, err)
		return
	}

	verses := strings.Split(song.Text, "\n\n")
	paginatedVerses := utils.PaginateSlice(verses, page, limit)
	if paginatedVerses == nil {
		utils.HandleError(w, "Страница не найдена", http.StatusNotFound, nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedVerses)
}

// DeleteSong godoc
// @Summary Удалить песню по ID
// @Description Удаляет песню из базы данных по её идентификатору
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} models.MessageResponse "Песня успешно удалена"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 404 {object} models.ErrorResponse "Песня не найдена"
// @Failure 500 {object} models.ErrorResponse "Ошибка при обработке запроса"
// @Router /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.HandleError(w, "Неверный ID песни", http.StatusBadRequest, err)
		return
	}

	var song models.Song
	if err := db.DB.First(&song, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.HandleError(w, "Песня не найдена", http.StatusNotFound, err)
			return
		}
		utils.HandleError(w, "Ошибка при поиске песни", http.StatusInternalServerError, err)
		return
	}

	if err := db.DB.Delete(&song).Error; err != nil {
		utils.HandleError(w, "Ошибка при удалении песни", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.MessageResponse{Message: "Песня успешно удалена"})
}

// UpdateSong godoc
// @Summary Изменить данные песни по ID
// @Description Обновляет информацию о песне по её идентификатору
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.UpdateSongRequest true "Данные для обновления песни"
// @Success 200 {object} models.Song "Обновленная песня"
// @Failure 400 {object} models.ErrorResponse "Неверный формат данных"
// @Failure 404 {object} models.ErrorResponse "Песня не найдена"
// @Failure 500 {object} models.ErrorResponse "Ошибка при обработке запроса"
// @Router /songs/{id} [put]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.HandleError(w, "Неверный ID песни", http.StatusBadRequest, err)
		return
	}

	var updateData models.UpdateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.HandleError(w, "Неверный формат данных", http.StatusBadRequest, err)
		return
	}

	var song models.Song
	if err := db.DB.First(&song, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.HandleError(w, "Песня не найдена", http.StatusNotFound, err)
			return
		}
		utils.HandleError(w, "Ошибка при поиске песни", http.StatusInternalServerError, err)
		return
	}

	// Обновление данных песни
	utils.ApplyUpdates(&song, updateData)

	if err := db.DB.Save(&song).Error; err != nil {
		utils.HandleError(w, "Ошибка при обновлении песни", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

// AddSong godoc
// @Summary Добавить новую песню
// @Description Добавляет новую песню и обогащает данные из внешнего API
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.AddSongRequest true "Данные для добавления песни"
// @Success 200 {object} models.Song "Добавленная песня"
// @Failure 400 {object} models.ErrorResponse "Неверный формат данных"
// @Failure 500 {object} models.ErrorResponse "Ошибка при обработке запроса"
// @Router /songs [post]
func AddSong(w http.ResponseWriter, r *http.Request) {
	// Парсим входящий запрос
	var addSongReq models.AddSongRequest
	if err := json.NewDecoder(r.Body).Decode(&addSongReq); err != nil {
		log.Printf("Ошибка при парсинге запроса: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Неверный формат данных"})
		return
	}

	// Проверяем наличие обязательных полей
	if addSongReq.Group == "" || addSongReq.Song == "" {
		log.Printf("Отсутствуют обязательные поля в запросе")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Поля 'group' и 'song' обязательны"})
		return
	}

	// Получаем данные песни через FetchSongDetails
	songDetail, err := utils.FetchSongDetails(addSongReq.Group, addSongReq.Song)
	if err != nil {
		log.Printf("Ошибка при запросе данных из внешнего API: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Ошибка внешнего сервиса"})
		return
	}

	// Создаем новую песню с обогащенными данными
	newSong := models.Song{
		MusicGroup:  addSongReq.Group,
		Song:        addSongReq.Song,
		ReleaseDate: songDetail.ReleaseDate,
		Text:        songDetail.Text,
		Link:        songDetail.Link,
	}

	// Сохраняем песню в базу данных
	if err := db.DB.Create(&newSong).Error; err != nil {
		log.Printf("Ошибка при сохранении песни в базу данных: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Ошибка при сохранении данных"})
		return
	}

	log.Printf("Песня '%s' группы '%s' успешно добавлена", newSong.Song, newSong.MusicGroup)

	// Возвращаем созданную песню в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSong)
}
