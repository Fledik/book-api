package main // Главный пакет приложения

import (
	"log" // Для вывода критических ошибок

	"github.com/Fledik/book-api/internal/database"   // Подключение SQLite
	"github.com/Fledik/book-api/internal/repository" // Репозитории
	"github.com/Fledik/book-api/internal/router"     // Роутер приложения
)

func main() {
	db, err := database.InitSQLite("data/books.db") // Открываем SQLite-базу
	if err != nil {
		log.Fatal(err) // Если база не открылась, завершаем приложение
	}
	defer db.Close() // Закрываем базу при завершении программы

	bookRepo := repository.NewBookRepository(db) // Создаём репозиторий книг

	r := router.SetupRouter(bookRepo) // Создаём роутер со всеми ручками

	r.Run(":8080") // Запускаем сервер на порту 8080
}