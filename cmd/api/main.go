package main // Точка входа приложения

import (
	"log" // Для вывода критических ошибок

	"github.com/Fledik/book-api/internal/database"   // SQLite
	"github.com/Fledik/book-api/internal/repository" // Репозиторий
	"github.com/Fledik/book-api/internal/router"     // Роутер
)

func main() {
	db, err := database.InitSQLite("data/books.db") // Подключаем базу
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Закрываем базу при остановке приложения

	bookRepo := repository.NewBookRepository(db) // Репозиторий книг

	r := router.SetupRouter(bookRepo) // Все HTTP-ручки

	if err := r.Run(":8080"); err != nil { // Запуск сервера
		log.Fatal(err)
	}
}