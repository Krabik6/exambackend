package handler

import (
	"exambackend/internal/model"
	"exambackend/pkg/jwtn"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) registerUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	id, err := h.userService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно зарегистрирован", "id": id})
}

// loginUser обрабатывает аутентификацию пользователя и выдачу токена
func (h *Handler) loginUser(c *gin.Context) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	user, err := h.userService.Authenticate(credentials.Login, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	token, err := jwtn.GenerateToken(user.ID, user.Role) // Предполагается, что функция generateToken уже реализована
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// createViolation обрабатывает создание нового заявления о нарушении
func (h *Handler) createViolation(c *gin.Context) {
	var violation model.Violation
	if err := c.BindJSON(&violation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	// Получение ID пользователя из контекста (предполагается, что мидлварь аутентификации уже добавлена и работает)
	userID, _ := getUserIDFromContext(c)

	violationID, err := h.violationService.CreateViolation(violation, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заявления о нарушении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заявление о нарушении успешно создано", "violationID": violationID})
}

// getViolationsByUser возвращает список заявлений, поданных пользователем
func (h *Handler) getViolationsByUser(c *gin.Context) {
	userID, _ := getUserIDFromContext(c) // Аналогично предыдущему методу

	violations, err := h.violationService.GetViolationsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка заявлений"})
		return
	}

	c.JSON(http.StatusOK, violations)
}

// updateViolationStatus изменяет статус заявления
func (h *Handler) updateViolationStatus(c *gin.Context) {
	violationID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var status struct {
		Status string `json:"status"`
	}
	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	err := h.violationService.UpdateViolationStatus(violationID, status.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса заявления"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус заявления успешно обновлен"})
}

// handler/handler.go
func (h *Handler) getAllViolations(c *gin.Context) {
	violations, err := h.violationService.GetAllViolations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка всех заявлений"})
		return
	}

	c.JSON(http.StatusOK, violations)
}
