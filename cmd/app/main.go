package main

import (
	"GO_WebApplication/internal/app/controllers"
	"GO_WebApplication/internal/app/repositories"
	"GO_WebApplication/internal/app/usecases"
	"database/sql"
	"log"
	"time"

	_ "GO_WebApplication/docs" // Импортируйте сгенерированные файлы документации
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// Logger - middleware для логирования
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("Запрос: %s %s, Статус: %d, Время обработки: %v\n",
			c.Request.Method,
			c.Request.URL,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

func main() {
	// Формируем строку подключения к базе данных из переменных окружения
	connStr := "user=postgres dbname=myappdatabase password=Admin123 host=localhost sslmode=disable"

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	// Инициализация репозиториев, use-case и контроллеров
	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	// Инициализация Gin
	router := gin.Default()

	// Подключаем middleware для логирования
	router.Use(Logger())

	// Обслуживание статических файлов (например, index.html)
	router.StaticFile("/", "./web/static/index.html")

	// Настройка маршрутов
	router.POST("/user", userController.CreateUser)
	router.GET("/users", userController.GetAllUsers)
	router.DELETE("/user", userController.DeleteUser)

	// Подключение Swagger UI (http://localhost:8080/swagger/index.html)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
