package repository // Пакет для работы с данными

import (
	"database/sql" // Стандартный SQL-пакет

	"github.com/Fledik/book-api/internal/models" // Модель книги
)

// BookRepository работает с книгами в базе данных
type BookRepository struct {
	db *sql.DB // Подключение к SQLite
}

// NewBookRepository создаёт новый репозиторий книг
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

// GetAll возвращает все книги из базы данных
func (r *BookRepository) GetAll() ([]models.Book, error) {
	query := `
	SELECT 
		id, isbn, title, author, description, publisher, year, genre, language, pages, rating, created_at, updated_at
	FROM books
	ORDER BY id DESC;`

	rows, err := r.db.Query(query) // Выполняем SELECT-запрос
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем rows после завершения функции

	books := make([]models.Book, 0) // Создаём пустой список книг

	for rows.Next() { // Проходим по каждой строке результата
		var book models.Book // В эту переменную считываем одну книгу

		err := rows.Scan(
			&book.ID,
			&book.ISBN,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Publisher,
			&book.Year,
			&book.Genre,
			&book.Language,
			&book.Pages,
			&book.Rating,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book) // Добавляем книгу в список
	}

	return books, nil
}