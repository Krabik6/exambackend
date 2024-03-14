package postgres

import (
	"database/sql"
	"errors"
	"exambackend/internal/model"
	"fmt"
)

// userRepo реализует UserRepository интерфейс
type userRepo struct {
	db *sql.DB
}

// NewUserRepo создает новый экземпляр userRepo
func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

// Create добавляет нового пользователя в базу данных
func (repo *userRepo) Create(user model.User) (int64, error) {
	query := `INSERT INTO users (login, password, full_name, phone, email) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var userID int64
	err := repo.db.QueryRow(query, user.Login, user.Password, user.FullName, user.Phone, user.Email).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при добавлении пользователя: %w", err)
	}
	return userID, nil
}

// FindByLogin возвращает пользователя по логину
func (repo *userRepo) FindByLogin(login string) (*model.User, error) {
	var user model.User
	// Добавляем выборку поля role в запрос
	query := `SELECT id, login, password, full_name, phone, email, role FROM users WHERE login = $1`
	err := repo.db.QueryRow(query, login).Scan(&user.ID, &user.Login, &user.Password, &user.FullName, &user.Phone, &user.Email, &user.Role) // Добавляем сюда &user.Role
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("ошибка при поиске пользователя: %w", err)
	}
	return &user, nil
}
