package main // Главный пакет приложения

import (
	"github.com/Fledik/book-api/internal/router" // Подключаем настройку маршрутов
)

func main() {
	r := router.SetupRouter() // Создаём роутер со всеми ручками

	r.Run(":8080") // Запускаем сервер на порту 8080
}