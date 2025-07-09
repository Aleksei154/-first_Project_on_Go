// Слой использования (Use Cases) содержит всю бизнес-логику web приложения
// вызывает методы их слоя репозитория

package usecases

import (
	"GO_WebApplication/internal/app/models"       // Импортируем модель пользователя
	"GO_WebApplication/internal/app/repositories" // Импортируем репозиторий пользователей
)

// UserUseCase представляет собой структуру, которая содержит ссылку на репозиторий пользователей.
type UserUseCase struct {
	repo *repositories.UserRepository // Поле repo хранит ссылку на репозиторий пользователей
}

// NewUserUseCase создает новый экземпляр UserUseCase.
// Он принимает указатель на UserRepository и возвращает указатель на UserUseCase.
func NewUserUseCase(repo *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo} // Возвращаем новый экземпляр UserUseCase с установленным репозиторием
}

// CreateUser создает нового пользователя, используя репозиторий.
// Он принимает указатель на модель пользователя и возвращает ошибку, если что-то пошло не так.
func (uc *UserUseCase) CreateUser(user *models.User) error {
	return uc.repo.CreateUser(user) // Вызываем метод CreateUser из репозитория
}

// GetAllUsers получает всех пользователей из репозитория.
// Он возвращает срез пользователей и ошибку, если что-то пошло не так.
func (uc *UserUseCase) GetAllUsers() ([]models.User, error) {
	return uc.repo.GetAllUsers() // Вызываем метод GetAllUsers из репозитория
}

// DeleteUser удаляет пользователя по его имени пользователя (username), используя репозиторий.
// Он принимает имя пользователя и возвращает ошибку, если что-то пошло не так.
func (uc *UserUseCase) DeleteUser(username string) error {
	return uc.repo.DeleteUser(username) // Вызываем метод DeleteUser из репозитория
}
