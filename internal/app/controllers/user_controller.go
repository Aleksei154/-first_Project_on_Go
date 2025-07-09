// Слой контроллеров (controllers) содержит логику обработки HTTP-запросов и ответов на них (например запросы пользователя)
// вызывает методы из слоя использования

package controllers

import (
	"GO_WebApplication/internal/app/models"   // Импортируем модель пользователя
	"GO_WebApplication/internal/app/usecases" // Импортируем слой использования
	"github.com/gin-gonic/gin"                // Импортируем Gin для работы с маршрутизацией
	"net/http"                                // Импортируем пакет для работы с HTTP
)

// UserController представляет контроллер для работы с пользователями.
type UserController struct {
	useCase *usecases.UserUseCase // Ссылка на слой использования для обработки бизнес-логики
}

// NewUserController создает новый экземпляр UserController.
func NewUserController(useCase *usecases.UserUseCase) *UserController {
	return &UserController{useCase: useCase} // Возвращаем новый экземпляр контроллера
}

// CreateUser обрабатывает HTTP-запрос на создание нового пользователя.
// @Summary Создает нового пользователя
// @Description Добавляет нового пользователя в базу данных
// @Accept json
// @Produce json
// @Param user body models.User true "Имя пользователя"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /user [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User // Создаем переменную для хранения данных пользователя
	// Привязываем JSON-данные из запроса к структуре user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // Возвращаем ошибку 400, если данные некорректны
		return
	}
	// Вызываем метод создания пользователя из слоя использования
	if err := uc.useCase.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) // Возвращаем ошибку 409, если пользователь уже существует
		return
	}
	c.JSON(http.StatusCreated, user) // Возвращаем статус 201 и данные созданного пользователя
}

// GetAllUsers обрабатывает HTTP-запрос на получение всех пользователей.
// @Summary Получает список всех пользователей
// @Description Возвращает массив пользователей
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	// Вызываем метод получения всех пользователей из слоя использования
	users, err := uc.useCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"}) // Возвращаем ошибку 500, если не удалось получить пользователей
		return
	}
	c.JSON(http.StatusOK, users) // Возвращаем статус 200 и список пользователей
}

// DeleteUser удаляет пользователя по имени пользователя.
// @Summary Удаляет пользователя
// @Description Удаляет пользователя из базы данных по имени
// @Param username query string true "Имя пользователя"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /user [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	username := c.Query("username") // Получаем имя пользователя из параметров запроса
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"}) // Возвращаем ошибку 400, если имя пользователя не указано
		return
	}
	// Вызываем метод удаления пользователя из слоя использования
	if err := uc.useCase.DeleteUser(username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Возвращаем ошибку 404, если пользователь не найден
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"}) // Возвращаем статус 200 и сообщение об успешном удалении
}
