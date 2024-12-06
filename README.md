# 🎶 Music Library API

**Music Library API** — это RESTful API для управления музыкальной библиотекой.
API предоставляет возможность добавления, редактирования, удаления, просмотра песен, а также получения их текста с пагинацией по куплетам.

## ✨ Возможности

- 🎧 **Получение списка песен** с фильтрацией по названию, группе, дате релиза и другим параметрам.
- ➕ **Добавление новой песни** с автоматическим обогащением данных из внешнего API.
- ✏️ **Обновление данных песни** по её идентификатору.
- ❌ **Удаление песни** из базы данных.
- 🔢 **Просмотр текста песни** с пагинацией по куплетам.

## 📚 Технологии

- **Go**: язык программирования.
- **Gorilla Mux**: маршрутизация и обработка HTTP-запросов.
- **GORM**: ORM для работы с базой данных.
- **PostgreSQL**: реляционная база данных.
- **Swagger**: документация и описание API.

## 🚀 Установка и запуск

1. 🔧 **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/ev1lyas/music-library.git
   cd music-library
   ```

2. 📝 **Установите зависимости:**
   ```bash
   go mod tidy
   ```

3. 🏢 **Настройте базу данных и внешний API в файле `music-library/.env`:**
   - Убедитесь, что PostgreSQL установлен и запущен.
   - Укажите параметры подключения в переменной окружения `.env`:
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=postgres
     DB_PASSWORD=yourpassword
     DB_NAME=music_library
     
     EXTERNAL_API_URL= URL
     SERVER_PORT= 8000
     ```

4. 🔄 **Предусмотрено заполнение тестовыми данными БД.** Для заполнения данными нужно раскомментировать код в `db/database.go`:
   ```bash
   /* log.Println("Заполнение тестовыми данными...")
   if err := ExecSQLFile(db, "migrations/002_insert_songs.sql"); err != nil {
       return fmt.Errorf("ошибка при заполнении тестовыми данными: %w", err)
   } */
   ```

5. 🔄 **Запустите сервер:**
   ```bash
   go run main.go
   ```

6. 🌐 **Откройте браузер и перейдите по адресу [http://localhost:8000/swagger/](http://localhost:8000/swagger/) для доступа к документации API.**

## 🔍 Примеры использования API

### 🎶 Получение списка песен
**GET** `/songs`

Параметры запроса (опционально):
- 🎤 `musicGroup`: название музыкальной группы.
- 🎵 `song`: название песни.
- 📅 `releaseDate`: дата релиза.
- 📝 `text`: текст песни.
- 🔗 `link`: ссылка на источник.
- 📄 `page`: номер страницы (по умолчанию 1).
- 🔢 `limit`: количество элементов на странице (по умолчанию 10).

Ответ:
```json
[
  {
    "id": 1,
    "musicGroup": "The Beatles",
    "song": "Yesterday",
    "releaseDate": "1965-08-14",
    "text": "Yesterday, all my troubles seemed so far away...",
    "link": "https://example.com/yesterday"
  }
]
```

### ➕ Добавление новой песни
**POST** `/songs`

Тело запроса:
```json
{
  "group": "The Beatles",
  "song": "Let It Be"
}
```

Ответ:
```json
{
  "id": 2,
  "musicGroup": "The Beatles",
  "song": "Let It Be",
  "releaseDate": "1970-03-06",
  "text": "When I find myself in times of trouble...",
  "link": "https://example.com/let-it-be"
}
```

### ✏️ Обновление данных песни
**PUT** `/songs/{id}`

Тело запроса:
```json
{
  "musicGroup": "The Beatles",
  "releaseDate": "1970-03-07"
}
```

Ответ:
```json
{
  "id": 2,
  "musicGroup": "The Beatles",
  "song": "Let It Be",
  "releaseDate": "1970-03-07",
  "text": "When I find myself in times of trouble...",
  "link": "https://example.com/let-it-be"
}
```

### ❌ Удаление песни
**DELETE** `/songs/{id}`

Ответ:
```json
{
  "message": "Песня успешно удалена"
}
```

## 🛠️ Средства разработки

- 🪵 **Логирование**: Включено логирование всех запросов через middleware.
- 🌐 **CORS**: Поддержка CORS для работы с фронтендом.
- 📃 **Swagger**: Полное описание API для интеграции и тестирования.
# music-library
