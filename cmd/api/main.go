package main // Главный пакет, с него запускается приложение

import (
	"net/http" // Стандартные HTTP-статусы

	"github.com/Fledik/book-api/internal/models" // Наша модель книги
	"github.com/gin-gonic/gin"                   // Gin — фреймворк для REST API
)

func main() {
	router := gin.Default() // Создаём роутер Gin

	router.GET("/health", func(c *gin.Context) { // Проверка, что сервер работает
		c.JSON(http.StatusOK, gin.H{ // JSON-ответ со статусом 200
			"status":  "ok",
			"service": "book-api",
		})
	})

	api := router.Group("/api/v1") // Группа маршрутов с версией API

	api.GET("/books", func(c *gin.Context) { // Получить список книг
		books := []models.Book{ // Пока временный список без базы данных
			{
				ID:          1,
				ISBN:        "9785171183661",
				Title:       "Мастер и Маргарита",
				Author:      "Михаил Булгаков",
				Description: "Роман о добре, зле, любви и свободе.",
				Publisher:   "АСТ",
				Year:        1967,
				Genre:       "Роман",
				Language:    "ru",
				Pages:       480,
				Rating:      4.8,
			},
			{
				ID:          2,
				ISBN:        "9785389177771",
				Title:       "1984",
				Author:      "Джордж Оруэлл",
				Description: "Антиутопия о тоталитарном обществе.",
				Publisher:   "Азбука",
				Year:        1949,
				Genre:       "Антиутопия",
				Language:    "ru",
				Pages:       320,
				Rating:      4.7,
			},
		}

		c.JSON(http.StatusOK, gin.H{ // Возвращаем список книг в JSON
			"items": books,      // Сами книги
			"count": len(books), // Количество книг
		})
	})

	router.Run(":8080") // Запуск сервера на порту 8080
}
