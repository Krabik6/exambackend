package handler

import (
	"exambackend/internal/middleware"
	"exambackend/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler структура для хранения ссылок на сервисы
type Handler struct {
	userService      service.UserService
	violationService service.ViolationService
}

// NewHandler создает экземпляр Handler с предоставленными сервисами
func NewHandler(userService service.UserService, violationService service.ViolationService) *Handler {
	return &Handler{
		userService:      userService,
		violationService: violationService,
	}
}

// RegisterRoutes регистрирует маршруты для обработки эндпоинтов
func (h *Handler) RegisterRoutes(router *gin.Engine) {
	// Публичные маршруты
	router.POST("/register", h.registerUser)
	router.POST("/login", h.loginUser)

	// Защищенные маршруты
	api := router.Group("/").Use(middleware.AuthMiddleware())
	{
		api.POST("/violations", h.createViolation)
		api.GET("/violations", h.getViolationsByUser)
	}

	// Маршруты для администратора
	admin := router.Group("/admin").Use(middleware.AuthMiddleware(), middleware.AdminRoleMiddleware())
	{
		admin.GET("/violations", h.getAllViolations)
		admin.PATCH("/violations/:id", h.updateViolationStatus) // Теперь только для админа
	}
}
