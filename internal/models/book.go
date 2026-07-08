package models // Пакет с моделями данных приложения

// Book описывает книгу в нашем API
type Book struct {
	ID          int64   `json:"id"`          // Уникальный ID книги
	ISBN        string  `json:"isbn"`        // Международный номер книги
	Title       string  `json:"title"`       // Название книги
	Author      string  `json:"author"`      // Автор книги
	Description string  `json:"description"` // Краткое описание
	Publisher   string  `json:"publisher"`   // Издательство
	Year        int     `json:"year"`        // Год публикации
	Genre       string  `json:"genre"`       // Жанр книги
	Language    string  `json:"language"`    // Язык книги
	Pages       int     `json:"pages"`       // Количество страниц
	Rating      float64 `json:"rating"`      // Рейтинг книги
	CreatedAt   string  `json:"created_at"`  // Дата создания записи
	UpdatedAt   string  `json:"updated_at"`  // Дата последнего обновления
}