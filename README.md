# Book API

REST API для книжного каталога на Go.

Проект реализует backend-приложение для работы с книгами: получение списка, поиск, добавление, изменение и удаление книг. Данные хранятся в SQLite, API работает в JSON-формате, приложение можно запускать локально или через Docker.

## Стек технологий

- Go
- Gin
- SQLite
- Docker
- Docker Compose
- REST API
- JSON

## Возможности

- Получение списка книг
- Получение книги по ID
- Добавление новой книги
- Изменение существующей книги
- Удаление книги
- Поиск по книгам
- Фильтрация по автору и жанру
- Пагинация через `limit` и `offset`
- Хранение данных в SQLite
- Запуск приложения в Docker
- Сохранение базы данных через volume-маунт папки `data`

## Структура проекта

```text
book-api/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── database/
│   │   └── sqlite.go
│   │
│   ├── handlers/
│   │   └── book_handler.go
│   │
│   ├── models/
│   │   └── book.go
│   │
│   ├── repository/
│   │   └── book_repository.go
│   │
│   └── router/
│       └── router.go
│
├── data/
│   └── .gitkeep
│
├── Dockerfile
├── docker-compose.yml
├── .dockerignore
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Модель книги

```json
{
  "id": 1,
  "isbn": "9785171183661",
  "title": "Мастер и Маргарита",
  "author": "Михаил Булгаков",
  "description": "Роман о добре, зле, любви и свободе.",
  "publisher": "АСТ",
  "year": 1967,
  "genre": "Роман",
  "language": "ru",
  "pages": 480,
  "rating": 4.8,
  "created_at": "2026-07-08 17:30:12",
  "updated_at": "2026-07-08 17:30:12"
}
```

## API endpoints

| Метод | URL | Описание |
|---|---|---|
| GET | `/health` | Проверка работы сервера |
| GET | `/api/v1/books` | Получить список книг |
| GET | `/api/v1/books/:id` | Получить книгу по ID |
| POST | `/api/v1/books` | Добавить новую книгу |
| PUT | `/api/v1/books/:id` | Изменить книгу |
| DELETE | `/api/v1/books/:id` | Удалить книгу |

## Query-параметры для списка книг

Ручка:

```text
GET /api/v1/books
```

Поддерживает параметры:

| Параметр | Описание |
|---|---|
| `search` | Поиск по названию, автору, описанию, издательству, жанру или ISBN |
| `author` | Фильтр по автору |
| `genre` | Фильтр по жанру |
| `limit` | Количество книг в ответе |
| `offset` | Смещение для пагинации |

Примеры:

```text
GET /api/v1/books?search=1984
GET /api/v1/books?author=Булгаков
GET /api/v1/books?genre=Роман
GET /api/v1/books?limit=10&offset=0
```

## Запуск локально

Перед запуском нужно установить зависимости:

```bash
go mod tidy
```

Запуск приложения:

```bash
go run ./cmd/api
```

После запуска сервер будет доступен по адресу:

```text
http://localhost:8080
```

Проверка:

```bash
Invoke-RestMethod http://localhost:8080/health
```

## Запуск через Docker

Сборка и запуск контейнера:

```bash
docker compose up --build
```

После запуска API будет доступно по адресу:

```text
http://localhost:8080
```

Остановка контейнера:

```bash
docker compose down
```

## Сохранение базы данных

SQLite-база создаётся автоматически в папке:

```text
data/books.db
```

В `docker-compose.yml` используется volume-маунт:

```yaml
volumes:
  - ./data:/app/data
```

Это значит, что папка `data` на компьютере подключается к папке `/app/data` внутри контейнера. Благодаря этому база данных сохраняется даже после остановки или пересоздания контейнера.

Файл базы данных не добавляется в GitHub, потому что он указан в `.gitignore`.

## Примеры запросов

### Проверка сервера

```powershell
Invoke-RestMethod http://localhost:8080/health
```

Пример ответа:

```text
service  status
-------  ------
book-api ok
```

### Получить список книг

```powershell
Invoke-RestMethod http://localhost:8080/api/v1/books
```

### Получить книгу по ID

```powershell
Invoke-RestMethod http://localhost:8080/api/v1/books/1
```

### Поиск книги

```powershell
Invoke-RestMethod "http://localhost:8080/api/v1/books?search=1984"
```

### Добавить книгу

```powershell
$book = Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/v1/books" -ContentType "application/json" -Body '{"isbn":"9780441172719","title":"Dune","author":"Frank Herbert","description":"Science fiction novel about politics, power and survival.","publisher":"Chilton Books","year":1965,"genre":"Science fiction","language":"en","pages":688,"rating":4.9}'

$book
```

### Изменить книгу

```powershell
Invoke-RestMethod -Method Put -Uri "http://localhost:8080/api/v1/books/$($book.id)" -ContentType "application/json" -Body '{"isbn":"9780441172719","title":"Dune Updated","author":"Frank Herbert","description":"Updated description.","publisher":"Chilton Books","year":1965,"genre":"Science fiction","language":"en","pages":700,"rating":5}'
```

### Удалить книгу

```powershell
Invoke-RestMethod -Method Delete -Uri "http://localhost:8080/api/v1/books/$($book.id)"
```

При успешном удалении возвращается статус `204 No Content`.

## Проверка работы

Были проверены основные сценарии:

- `/health` возвращает статус сервиса
- `/api/v1/books` возвращает список книг
- `/api/v1/books/1` возвращает книгу по ID
- поиск через `search` работает
- `POST` создаёт новую книгу
- `PUT` изменяет книгу
- `DELETE` удаляет книгу

## Назначение проекта

Проект создан в рамках практики для изучения разработки backend-приложений на Go. В проекте показана работа с REST API, обработчиками Gin, SQLite-базой данных, Docker-контейнеризацией и сохранением данных через volume-маунт.