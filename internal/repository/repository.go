package repository

import "exambackend/internal/model"

// UserRepository определяет интерфейс для работы с пользователями
type UserRepository interface {
	Create(user model.User) (int64, error)
	FindByLogin(login string) (*model.User, error)
	// Дополнительные методы
}
