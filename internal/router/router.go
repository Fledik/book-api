package router // Пакет для настройки маршрутов приложения

import (
	"net/http" // HTTP-статусы

	"github.com/Fledik/book-api/internal/handlers" // Наши обработчики
	"github.com/gin-gonic/gin"                    // Gin-роутер
)

// SetupRouter создаёт и настраивает все маршруты API
func SetupRouter() *gin.Engine {
	router := gin.Default() // Создаём Gin-роутер

	router.GET("/health", func(c *gin.Context) { // Проверка работы сервера
		c.JSON(http.StatusOK, gin.H{ // Возвращаем JSON со статусом 200
			"status":  "ok",
			"service": "book-api",
		})
	})

	bookHandler := handlers.NewBookHandler() // Создаём обработчик книг

	api := router.Group("/api/v1") // Группа маршрутов версии v1
	{
		api.GET("/books", bookHandler.GetBooks) // Получить список книг
	}

	return router // Возвращаем готовый роутер
}