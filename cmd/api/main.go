package main // Главный пакет приложения, нужен для запуска программы

import (
	"net/http" // Стандартный пакет Go для HTTP-статусов

	"github.com/gin-gonic/gin" // Gin — фреймворк для создания REST API
)

func main() {
	router := gin.Default() // Создаём роутер Gin с логированием и обработкой ошибок

	router.GET("/health", func(c *gin.Context) { // GET-ручка для проверки работы сервера
		c.JSON(http.StatusOK, gin.H{ // Отправляем JSON-ответ со статусом 200
			"status":  "ok",       // Статус приложения
			"service": "book-api", // Название сервиса
		})
	})

	router.Run(":8080") // Запускаем сервер на порту 8080
}