package service

import (
	"exambackend/internal/model"
	"exambackend/internal/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// UserService предоставляет методы для работы с пользователями
type UserService interface {
	Register(user model.User) (int64, error)
	Authenticate(login, password string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService создает новый экземпляр UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Register регистрирует нового пользователя
func (s *userService) Register(user model.User) (int64, error) {
	// Хеширование пароля пользователя перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("не удалось хешировать пароль: %w", err)
	}
	user.Password = string(hashedPassword) // Заменяем исходный пароль на его хеш
	fmt.Println("user.Password:", user.Password)
	// Здесь могут быть дополнительные проверки и логика перед добавлением пользователя
	return s.userRepo.Create(user)
}

// Authenticate проверяет логин и пароль пользователя
func (s *userService) Authenticate(login, password string) (*model.User, error) {
	user, err := s.userRepo.FindByLogin(login)
	if err != nil {
		return nil, err
	}
	if user == nil || !checkPasswordHash(password, user.Password) {
		// checkPasswordHash - это функция для проверки пароля (не показана здесь)
		return nil, fmt.Errorf("неверный логин или пароль")
	}
	return user, nil
}
