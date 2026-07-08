package models // Модели данных приложения

// Book — основная структура книги, которая отдаётся клиенту
type Book struct {
	ID          int64   `json:"id"`          // Уникальный ID книги
	ISBN        string  `json:"isbn"`        // ISBN книги
	Title       string  `json:"title"`       // Название
	Author      string  `json:"author"`      // Автор
	Description string  `json:"description"` // Описание
	Publisher   string  `json:"publisher"`   // Издательство
	Year        int     `json:"year"`        // Год выпуска
	Genre       string  `json:"genre"`       // Жанр
	Language    string  `json:"language"`    // Язык
	Pages       int     `json:"pages"`       // Количество страниц
	Rating      float64 `json:"rating"`      // Рейтинг от 0 до 5
	CreatedAt   string  `json:"created_at"`  // Дата создания записи
	UpdatedAt   string  `json:"updated_at"`  // Дата обновления записи
}

// BookInput — структура для создания и изменения книги
type BookInput struct {
	ISBN        string  `json:"isbn"`        // ISBN можно не указывать
	Title       string  `json:"title"`       // Название обязательно
	Author      string  `json:"author"`      // Автор обязателен
	Description string  `json:"description"` // Описание
	Publisher   string  `json:"publisher"`   // Издательство
	Year        int     `json:"year"`        // Год выпуска
	Genre       string  `json:"genre"`       // Жанр
	Language    string  `json:"language"`    // Язык
	Pages       int     `json:"pages"`       // Количество страниц
	Rating      float64 `json:"rating"`      // Рейтинг
}

// ListBooksParams — параметры для списка книг, поиска и фильтрации
type ListBooksParams struct {
	Search string // Поиск по названию, автору, описанию, жанру, ISBN
	Author string // Фильтр по автору
	Genre  string // Фильтр по жанру
	Limit  int    // Сколько книг вернуть
	Offset int    // Сколько книг пропустить
}