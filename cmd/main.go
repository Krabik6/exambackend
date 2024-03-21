package main

import (
	"exambackend/internal/config"
	"exambackend/internal/handler"
	"exambackend/internal/repository/postgres"
	"exambackend/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	cfg := config.New()
	fmt.Println(cfg)
	// Инициализация подключения к БД
	db, err := postgres.NewDB(cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.User, cfg.DBConfig.Password, cfg.DBConfig.DBName)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	// Создание репозиториев
	userRepo := postgres.NewUserRepo(db)
	violationRepo := postgres.NewViolationRepo(db)

	// Создание сервисов
	userService := service.NewUserService(userRepo)
	violationService := service.NewViolationService(violationRepo)

	// Создание хендлера и регистрация маршрутов
	h := handler.NewHandler(userService, violationService)
	router := gin.Default()
	router.Use(handler.CORSMiddleware())

	h.RegisterRoutes(router)

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
