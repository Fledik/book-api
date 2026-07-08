package handlers // Пакет с HTTP-обработчиками

import (
	"net/http" // HTTP-статусы, например 200 OK

	"github.com/Fledik/book-api/internal/models" // Модель книги
	"github.com/gin-gonic/gin"                  // Gin-контекст для обработки запросов
)

// BookHandler хранит методы для работы с книгами
type BookHandler struct{}

// NewBookHandler создаёт новый обработчик книг
func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

// GetBooks возвращает список книг
func (h *BookHandler) GetBooks(c *gin.Context) {
	books := []models.Book{ // Временные данные, позже заменим на SQLite
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

	c.JSON(http.StatusOK, gin.H{ // Возвращаем JSON-ответ
		"items": books,      // Список книг
		"count": len(books), // Количество книг
	})
}