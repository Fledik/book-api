package database // Работа с базой данных

import (
	"database/sql" // Стандартный пакет для SQL
	"fmt"          // Форматирование ошибок
	"os"           // Создание папок
	"path/filepath" // Работа с путями файлов

	_ "modernc.org/sqlite" // SQLite-драйвер
)

// InitSQLite открывает SQLite и подготавливает таблицу
func InitSQLite(dbPath string) (*sql.DB, error) {
	dataDir := filepath.Dir(dbPath) // Получаем папку, где будет лежать база

	err := os.MkdirAll(dataDir, os.ModePerm) // Создаём папку data, если её нет
	if err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath) // Открываем SQLite-файл
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = createBooksTable(db) // Создаём таблицу books
	if err != nil {
		return nil, err
	}

	err = seedBooks(db) // Добавляем стартовые книги
	if err != nil {
		return nil, err
	}

	return db, nil // Возвращаем подключение к базе
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

	_, err := db.Exec(query) // Выполняем SQL
	if err != nil {
		return fmt.Errorf("failed to create books table: %w", err)
	}

	return nil
}

// seedBooks добавляет тестовые книги, если база пустая
func seedBooks(db *sql.DB) error {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count) // Считаем книги
	if err != nil {
		return fmt.Errorf("failed to count books: %w", err)
	}

	if count > 0 {
		return nil // Если книги уже есть, ничего не добавляем
	}

	query := `
	INSERT OR IGNORE INTO books
	(isbn, title, author, description, publisher, year, genre, language, pages, rating)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

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