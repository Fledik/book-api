package handlers // HTTP-обработчики

import (
	"net/http" // HTTP-статусы
	"strconv"  // Перевод строки в число
	"strings"  // Проверка пустых строк

	"github.com/Fledik/book-api/internal/models"     // Модели
	"github.com/Fledik/book-api/internal/repository" // Репозиторий
	"github.com/gin-gonic/gin"                       // Gin
)

// BookHandler хранит зависимости для ручек книг
type BookHandler struct {
	repo *repository.BookRepository // Работа с книгами в БД
}

// NewBookHandler создаёт обработчик книг
func NewBookHandler(repo *repository.BookRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

// GetBooks возвращает список книг
func (h *BookHandler) GetBooks(c *gin.Context) {
	limit := parseIntQuery(c, "limit", 20)   // Сколько книг вернуть
	offset := parseIntQuery(c, "offset", 0)  // Сколько книг пропустить

	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100 // Ограничение, чтобы не отдавать слишком много
	}
	if offset < 0 {
		offset = 0
	}

	params := models.ListBooksParams{
		Search: c.Query("search"), // /books?search=...
		Author: c.Query("author"), // /books?author=...
		Genre:  c.Query("genre"),  // /books?genre=...
		Limit:  limit,
		Offset: offset,
	}

	books, err := h.repo.List(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items":  books,
		"count":  len(books),
		"limit":  limit,
		"offset": offset,
	})
}

// GetBookByID возвращает книгу по ID
func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	book, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get book"})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook создаёт новую книгу
func (h *BookHandler) CreateBook(c *gin.Context) {
	var input models.BookInput

	err := c.ShouldBindJSON(&input) // Читаем JSON из тела запроса
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	if message := validateBookInput(input); message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	book, err := h.repo.Create(input)
	if err != nil {
		if isUniqueError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "book with this ISBN already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// UpdateBook изменяет книгу по ID
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var input models.BookInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	if message := validateBookInput(input); message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	book, err := h.repo.Update(id, input)
	if err != nil {
		if isUniqueError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "book with this ISBN already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update book"})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook удаляет книгу по ID
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	deleted, err := h.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete book"})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.Status(http.StatusNoContent) // 204 без тела ответа
}

// parseID достаёт ID из URL
func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return 0, false
	}

	return id, true
}

// parseIntQuery читает числовой query-параметр
func parseIntQuery(c *gin.Context, name string, defaultValue int) int {
	value := c.Query(name)
	if value == "" {
		return defaultValue
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}

// validateBookInput проверяет данные книги
func validateBookInput(input models.BookInput) string {
	if strings.TrimSpace(input.Title) == "" {
		return "title is required"
	}

	if strings.TrimSpace(input.Author) == "" {
		return "author is required"
	}

	if input.Pages < 0 {
		return "pages cannot be negative"
	}

	if input.Rating < 0 || input.Rating > 5 {
		return "rating must be between 0 and 5"
	}

	if input.Year < 0 {
		return "year cannot be negative"
	}

	return ""
}

// isUniqueError проверяет ошибку уникальности ISBN
func isUniqueError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "unique")
}