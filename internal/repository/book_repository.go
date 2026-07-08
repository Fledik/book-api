package repository // Слой работы с данными

import (
	"database/sql" // SQL-пакет
	"strings"      // Для сборки условий поиска

	"github.com/Fledik/book-api/internal/models" // Модели приложения
)

// BookRepository работает с таблицей books
type BookRepository struct {
	db *sql.DB // Подключение к SQLite
}

// NewBookRepository создаёт репозиторий книг
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

// scanner нужен, чтобы одинаково сканировать QueryRow и Rows
type scanner interface {
	Scan(dest ...any) error
}

// Общий набор полей для SELECT-запросов
const bookSelectColumns = `
	id,
	COALESCE(isbn, ''),
	COALESCE(title, ''),
	COALESCE(author, ''),
	COALESCE(description, ''),
	COALESCE(publisher, ''),
	COALESCE(year, 0),
	COALESCE(genre, ''),
	COALESCE(language, ''),
	COALESCE(pages, 0),
	COALESCE(rating, 0),
	COALESCE(created_at, ''),
	COALESCE(updated_at, '')
`

// scanBook превращает строку из базы в структуру Book
func scanBook(s scanner) (*models.Book, error) {
	var book models.Book

	err := s.Scan(
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

	return &book, nil
}

// nullIfEmpty сохраняет пустой ISBN как NULL, чтобы не ломать UNIQUE
func nullIfEmpty(value string) any {
	if strings.TrimSpace(value) == "" {
		return nil
	}

	return value
}

// List возвращает список книг с поиском, фильтрами и пагинацией
func (r *BookRepository) List(params models.ListBooksParams) ([]models.Book, error) {
	query := "SELECT " + bookSelectColumns + " FROM books"

	conditions := make([]string, 0) // WHERE-условия
	args := make([]any, 0)           // Значения для SQL-параметров

	if strings.TrimSpace(params.Search) != "" {
		search := "%" + strings.TrimSpace(params.Search) + "%"

		conditions = append(conditions, `
			(title LIKE ? OR author LIKE ? OR description LIKE ? OR publisher LIKE ? OR genre LIKE ? OR isbn LIKE ?)
		`)

		args = append(args, search, search, search, search, search, search)
	}

	if strings.TrimSpace(params.Author) != "" {
		conditions = append(conditions, "author LIKE ?")
		args = append(args, "%"+strings.TrimSpace(params.Author)+"%")
	}

	if strings.TrimSpace(params.Genre) != "" {
		conditions = append(conditions, "genre LIKE ?")
		args = append(args, "%"+strings.TrimSpace(params.Genre)+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, params.Limit, params.Offset)

	rows, err := r.db.Query(query, args...) // Выполняем SELECT
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]models.Book, 0)

	for rows.Next() {
		book, err := scanBook(rows)
		if err != nil {
			return nil, err
		}

		books = append(books, *book)
	}

	return books, nil
}

// GetByID возвращает одну книгу по ID
func (r *BookRepository) GetByID(id int64) (*models.Book, error) {
	query := "SELECT " + bookSelectColumns + " FROM books WHERE id = ?"

	book, err := scanBook(r.db.QueryRow(query, id))
	if err == sql.ErrNoRows {
		return nil, nil // Книга не найдена
	}
	if err != nil {
		return nil, err
	}

	return book, nil
}

// Create добавляет новую книгу
func (r *BookRepository) Create(input models.BookInput) (*models.Book, error) {
	if strings.TrimSpace(input.Language) == "" {
		input.Language = "ru" // Значение по умолчанию
	}

	query := `
	INSERT INTO books
	(isbn, title, author, description, publisher, year, genre, language, pages, rating)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	result, err := r.db.Exec(
		query,
		nullIfEmpty(input.ISBN),
		input.Title,
		input.Author,
		input.Description,
		input.Publisher,
		input.Year,
		input.Genre,
		input.Language,
		input.Pages,
		input.Rating,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId() // Получаем ID новой книги
	if err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

// Update изменяет книгу по ID
func (r *BookRepository) Update(id int64, input models.BookInput) (*models.Book, error) {
	if strings.TrimSpace(input.Language) == "" {
		input.Language = "ru"
	}

	query := `
	UPDATE books
	SET 
		isbn = ?,
		title = ?,
		author = ?,
		description = ?,
		publisher = ?,
		year = ?,
		genre = ?,
		language = ?,
		pages = ?,
		rating = ?,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = ?;`

	result, err := r.db.Exec(
		query,
		nullIfEmpty(input.ISBN),
		input.Title,
		input.Author,
		input.Description,
		input.Publisher,
		input.Year,
		input.Genre,
		input.Language,
		input.Pages,
		input.Rating,
		id,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, nil // Книга не найдена
	}

	return r.GetByID(id)
}

// Delete удаляет книгу по ID
func (r *BookRepository) Delete(id int64) (bool, error) {
	result, err := r.db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected > 0, nil
}