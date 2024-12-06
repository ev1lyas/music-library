package models

// Song представляет модель песни
type Song struct {
	ID          uint   `gorm:"primaryKey;column:id" json:"id"`
	MusicGroup  string `gorm:"column:musicgroup" json:"musicGroup"`
	Song        string `gorm:"column:song" json:"song"`
	ReleaseDate string `gorm:"column:releasedate" json:"releaseDate"`
	Text        string `gorm:"column:text" json:"text"`
	Link        string `gorm:"column:link" json:"link"`
}

// ErrorResponse представляет структуру ошибки
type ErrorResponse struct {
	Error string `json:"error"`
}

// VerseResponse представляет ответ с куплетами песни
type VerseResponse struct {
	SongID     uint     `json:"song_id"`
	MusicGroup string   `json:"musicGroup"`
	SongName   string   `json:"song_name"`
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
	Total      int      `json:"total_verses"`
	Verses     []string `json:"verses"`
}

// MessageResponse представляет ответ с сообщением
type MessageResponse struct {
	Message string `json:"message"`
}

// UpdateSongRequest представляет данные для обновления песни
type UpdateSongRequest struct {
	MusicGroup  *string `json:"musicGroup,omitempty"`
	Song        *string `json:"song,omitempty"`
	ReleaseDate *string `json:"releaseDate,omitempty"`
	Text        *string `json:"text,omitempty"`
	Link        *string `json:"link,omitempty"`
}

// Cтруктура для представления входящего запроса при добавлении песни
type AddSongRequest struct {
	Group string `json:"group" validate:"required"`
	Song  string `json:"song" validate:"required"`
}

// Cтруктура для представления ответа от внешнего API
type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
