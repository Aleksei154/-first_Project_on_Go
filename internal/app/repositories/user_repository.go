// Слой репозиториев (repositories) содержит всю логику взаимодействия с базой данных
//(подключение и методы работы)

package repositories

import (
	"GO_WebApplication/internal/app/models" // Импортируем модель пользователя
	"database/sql"                          // Импортируем пакет для работы с базой данных
	"fmt"                                   // Импортируем пакет для форматирования строк

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

// UserRepository представляет собой структуру, которая содержит подключение и методы работы с базой данных.
type UserRepository struct {
	db *sql.DB // Поле db хранит соединение с базой данных
}

// NewUserRepository создает новый экземпляр UserRepository.
// Он принимает указатель на sql.DB и возвращает указатель на UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db} // Возвращаем новый экземпляр UserRepository с установленным соединением
}

// CreateUser добавляет нового пользователя в базу данных.
func (r *UserRepository) CreateUser(user *models.User) error {
	var exists bool // Переменная для проверки существования пользователя

	// Проверяем, существует ли уже пользователь с таким именем
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", user.Username).Scan(&exists)
	if err != nil {
		return err // Возвращаем ошибку, если произошла ошибка при выполнении запроса
	}

	// Если пользователь с таким именем уже существует, возвращаем ошибку
	if exists {
		return fmt.Errorf("username %s already exists", user.Username)
	}

	// Если пользователь не существует, добавляем его в базу данных
	_, err = r.db.Exec("INSERT INTO users (username) VALUES ($1)", user.Username)
	return err // Возвращаем ошибку, если произошла ошибка при добавлении пользователя
}

// GetAllUsers извлекает всех пользователей из базы данных.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	// Выполняем запрос для получения всех пользователей
	rows, err := r.db.Query("SELECT id, username FROM users")
	if err != nil {
		return nil, err // Возвращаем nil и ошибку, если произошла ошибка при выполнении запроса
	}
	defer rows.Close() // Закрываем rows после завершения работы с ними

	var users []models.User // Создаем срез для хранения пользователей

	// Проходим по всем строкам результата запроса
	for rows.Next() {
		var user models.User // Создаем переменную для хранения данных пользователя
		// Сканируем данные из строки в переменную user
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err // Возвращаем nil и ошибку, если произошла ошибка при сканировании
		}
		users = append(users, user) // Добавляем пользователя в срез
	}

	return users, nil // Возвращаем срез пользователей и nil (без ошибок)
}

// DeleteUser удаляет пользователя из базы данных по его имени пользователя (username).
// Он принимает имя пользователя и возвращает ошибку, если что-то пошло не так.
func (r *UserRepository) DeleteUser(username string) error {
	// Выполняем запрос для удаления пользователя с указанным именем
	result, err := r.db.Exec("DELETE FROM users WHERE username = $1", username)
	if err != nil {
		return err // Возвращаем ошибку, если произошла ошибка при выполнении запроса
	}

	// Проверяем, было ли удалено хотя бы одно значение
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err // Возвращаем ошибку, если произошла ошибка при получении количества затронутых строк
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with username %s does not exist", username) // Возвращаем ошибку, если пользователь не найден
	}

	return nil // Возвращаем nil, если удаление прошло успешно
}
