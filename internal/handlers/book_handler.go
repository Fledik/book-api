package handlers // Пакет с HTTP-обработчиками

import (
	"net/http" // HTTP-статусы

	"github.com/Fledik/book-api/internal/repository" // Репозиторий книг
	"github.com/gin-gonic/gin"                      // Gin-контекст
)

// BookHandler хранит зависимости для ручек книг
type BookHandler struct {
	repo *repository.BookRepository // Репозиторий для работы с книгами
}

// NewBookHandler создаёт новый обработчик книг
func NewBookHandler(repo *repository.BookRepository) *BookHandler {
	return &BookHandler{
		repo: repo,
	}
}

// GetBooks возвращает список книг из SQLite
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.repo.GetAll() // Получаем книги из базы данных
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ // Если ошибка БД, отдаём 500
			"error": "failed to get books",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{ // Возвращаем JSON-ответ
		"items": books,      // Список книг
		"count": len(books), // Количество книг
	})
}