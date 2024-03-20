package model

// User представляет пользователя системы
type User struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"` // Добавляем поле для роли пользователя
}
