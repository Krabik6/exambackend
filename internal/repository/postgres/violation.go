package postgres

import (
	"database/sql"
	"exambackend/internal/model"
	"fmt"
)

// violationRepo реализует ViolationRepository интерфейс
type violationRepo struct {
	db *sql.DB
}

// NewViolationRepo создает новый экземпляр violationRepo
func NewViolationRepo(db *sql.DB) *violationRepo {
	return &violationRepo{db: db}
}

func (r *violationRepo) GetAllViolations() ([]model.Violation, error) {
	var violations []model.Violation
	query := `SELECT id, user_id, car_number, description, status, created_at FROM violations`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying all violations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v model.Violation
		if err := rows.Scan(&v.ID, &v.UserID, &v.CarNumber, &v.Description, &v.Status, &v.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning violation: %w", err)
		}
		violations = append(violations, v)
	}

	return violations, nil
}

// Create добавляет новое заявление о нарушении в базу данных
func (repo *violationRepo) Create(violation model.Violation) (int64, error) {
	query := `INSERT INTO violations (user_id, car_number, description) VALUES ($1, $2, $3) RETURNING id`
	var violationID int64
	err := repo.db.QueryRow(query, violation.UserID, violation.CarNumber, violation.Description).Scan(&violationID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при добавлении заявления о нарушении: %w", err)
	}
	return violationID, nil
}

func (repo *violationRepo) FindByUserID(userID int64) ([]model.Violation, error) {
	var violations []model.Violation

	query := `SELECT id, user_id, car_number, description, status, created_at FROM violations WHERE user_id = $1`
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе заявлений: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v model.Violation
		if err := rows.Scan(&v.ID, &v.UserID, &v.CarNumber, &v.Description, &v.Status, &v.CreatedAt); err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки заявлений: %w", err)
		}
		violations = append(violations, v)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка после обработки всех строк заявлений: %w", err)
	}

	return violations, nil
}

func (repo *violationRepo) UpdateStatus(violationID int64, status string) error {
	query := `UPDATE violations SET status = $1 WHERE id = $2`
	res, err := repo.db.Exec(query, status, violationID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении статуса заявления: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при получении количества обновленных записей: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("заявление не найдено")
	}

	return nil
}
