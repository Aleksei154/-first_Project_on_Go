// package repositories
//
// import (
//
//	"GO_WebApplication/internal/app/models"
//	"database/sql"
//	_ "github.com/lib/pq" // PostgreSQL driver
//	"testing"
//
// )
//
//	func TestCreateUser(t *testing.T) {
//		// Настройка тестовой базы данных
//		db, err := sql.Open("postgres", "user=postgres dbname=myappdatabase password=Admin123 host=localhost sslmode=disable")
//		if err != nil {
//			t.Fatal(err)
//		}
//		defer db.Close()
//
//		// Очистка таблицы пользователей
//		_, err = db.Exec("DELETE FROM users")
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		repo := NewUserRepository(db)
//		user := &models.User{Username: "testuser"}
//
//		// Тестируем создание пользователя
//		err = repo.CreateUser(user)
//		if err != nil {
//			t.Fatalf("expected no error, got %v", err)
//		}
//
//		// Проверяем, что пользователь был создан
//		users, err := repo.GetAllUsers()
//		if err != nil || len(users) == 0 {
//			t.Fatalf("expected to find users, got %v", err)
//		}
//	}
package repositories

import (
	"GO_WebApplication/internal/app/models"
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Настройка тестовой базы данных
	db, err := sql.Open("postgres", "user=postgres dbname=myappdatabase password=Admin123 host=localhost sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Очистка таблицы пользователей
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)
	user := &models.User{Username: "testuser"}

	// Тестируем создание пользователя
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Проверяем, что пользователь был создан
	users, err := repo.GetAllUsers()
	if err != nil || len(users) == 0 {
		t.Fatalf("expected to find users, got %v", err)
	}
}

func TestGetAllUsers(t *testing.T) {
	// Настройка тестовой базы данных
	db, err := sql.Open("postgres", "user=postgres dbname=myappdatabase password=Admin123 host=localhost sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Очистка таблицы пользователей
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	// Добавляем тестового пользователя
	user := &models.User{Username: "testuser"}
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Тестируем получение всех пользователей
	users, err := repo.GetAllUsers()
	if err != nil || len(users) == 0 {
		t.Fatalf("expected to find users, got %v", err)
	}

	if len(users) != 1 || users[0].Username != "testuser" {
		t.Fatalf("expected to find one user with username 'testuser', got %v", users)
	}
}

func TestDeleteUser(t *testing.T) {
	// Настройка тестовой базы данных
	db, err := sql.Open("postgres", "user=postgres dbname=myappdatabase password=Admin123 host=localhost sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Очистка таблицы пользователей
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	// Добавляем тестового пользователя
	user := &models.User{Username: "testuser"}
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Тестируем удаление пользователя
	err = repo.DeleteUser("testuser")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Проверяем, что пользователь был удален
	users, err := repo.GetAllUsers()
	if err != nil || len(users) != 0 {
		t.Fatalf("expected no users, got %v", users)
	}
}

func TestDeleteNonExistentUser(t *testing.T) {
	// Настройка тестовой базы данных
	db, err := sql.Open("postgres", "user=postgres dbname=myappdatabase password=Admin123 host=localhost sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Очистка таблицы пользователей
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	// Тестируем удаление несуществующего пользователя
	err = repo.DeleteUser("nonexistentuser")
	if err == nil {
		t.Fatal("expected error when deleting nonexistent user, got nil")
	}
}
func TestCreateUserWithExistingUsername(t *testing.T) {
	// Настройка тестовой базы данных
	db, err := sql.Open("postgres", "user=postgres dbname=myappdatabase password=Admin123 host=localhost sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Очистка таблицы пользователей
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	// Создаем первого пользователя
	user1 := &models.User{Username: "testuser"}
	err = repo.CreateUser(user1)
	if err != nil {
		t.Fatalf("expected no error when creating first user, got %v", err)
	}

	// Пытаемся создать второго пользователя с тем же именем
	user2 := &models.User{Username: "testuser"}
	err = repo.CreateUser(user2)

	// Проверяем, что возникла ошибка
	if err == nil {
		t.Fatal("expected error when creating user with existing username, got nil")
	}

	// Проверяем, что сообщение об ошибке соответствует ожидаемому
	expectedError := "username testuser already exists"
	if err.Error() != expectedError {
		t.Fatalf("expected error message '%s', got '%s'", expectedError, err.Error())
	}
}
