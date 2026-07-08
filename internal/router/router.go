package router // Настройка маршрутов

import (
	"net/http" // HTTP-статусы

	"github.com/Fledik/book-api/internal/handlers"   // HTTP-обработчики
	"github.com/Fledik/book-api/internal/repository" // Репозитории
	"github.com/gin-gonic/gin"                       // Gin
)

// SetupRouter создаёт все маршруты приложения
func SetupRouter(bookRepo *repository.BookRepository) *gin.Engine {
	router := gin.Default() // Роутер Gin с логами

	router.GET("/health", func(c *gin.Context) { // Проверка работы сервера
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "book-api",
		})
	})

	bookHandler := handlers.NewBookHandler(bookRepo) // Ручки книг

	api := router.Group("/api/v1") // Версия API
	{
		api.GET("/books", bookHandler.GetBooks)          // Список + поиск
		api.GET("/books/:id", bookHandler.GetBookByID)   // Одна книга
		api.POST("/books", bookHandler.CreateBook)       // Добавить книгу
		api.PUT("/books/:id", bookHandler.UpdateBook)    // Изменить книгу
		api.DELETE("/books/:id", bookHandler.DeleteBook) // Удалить книгу
	}

	return router
}