package database // Пакет для работы с базой данных

import (
	"database/sql" // Стандартный пакет Go для работы с SQL-базами
	"fmt"          // Нужен для форматирования ошибок
	"os"           // Нужен для создания папки data

	_ "modernc.org/sqlite" // SQLite-драйвер, подключается через blank import
)

// InitSQLite открывает базу данных и создаёт таблицу books
func InitSQLite(dbPath string) (*sql.DB, error) {
	err := os.MkdirAll("data", os.ModePerm) // Создаём папку data, если её нет
	if err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath) // Открываем SQLite-базу по пути
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = createBooksTable(db) // Создаём таблицу books, если её ещё нет
	if err != nil {
		return nil, err
	}

	err = seedBooks(db) // Добавляем тестовые книги, если таблица пустая
	if err != nil {
		return nil, err
	}

	return db, nil // Возвращаем готовое подключение к базе
}

// createBooksTable создаёт таблицу книг
func createBooksTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		isbn TEXT UNIQUE,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		description TEXT,
		publisher TEXT,
		year INTEGER,
		genre TEXT,
		language TEXT DEFAULT 'ru',
		pages INTEGER,
		rating REAL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query) // Выполняем SQL-запрос создания таблицы
	if err != nil {
		return fmt.Errorf("failed to create books table: %w", err)
	}

	return nil
}

// seedBooks добавляет стартовые книги, чтобы API сразу что-то возвращал
func seedBooks(db *sql.DB) error {
	var count int // Сюда запишем количество книг в таблице

	err := db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count) // Считаем книги
	if err != nil {
		return fmt.Errorf("failed to count books: %w", err)
	}

	if count > 0 {
		return nil // Если книги уже есть, повторно ничего не добавляем
	}

	query := `
	INSERT INTO books 
	(isbn, title, author, description, publisher, year, genre, language, pages, rating)
	VALUES 
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err = db.Exec(
		query,
		"9785171183661",
		"Мастер и Маргарита",
		"Михаил Булгаков",
		"Роман о добре, зле, любви и свободе.",
		"АСТ",
		1967,
		"Роман",
		"ru",
		480,
		4.8,
	)
	if err != nil {
		return fmt.Errorf("failed to seed first book: %w", err)
	}

	_, err = db.Exec(
		query,
		"9785389177771",
		"1984",
		"Джордж Оруэлл",
		"Антиутопия о тоталитарном обществе.",
		"Азбука",
		1949,
		"Антиутопия",
		"ru",
		320,
		4.7,
	)
	if err != nil {
		return fmt.Errorf("failed to seed second book: %w", err)
	}

	return nil
}