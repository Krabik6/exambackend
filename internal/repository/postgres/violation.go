package postgres

import (
	"database/sql"
	"exambackend/internal/model"
	"fmt"
	"github.com/lib/pq"
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
	query := `SELECT v.id, v.user_id, v.car_number, v.description, v.status, v.created_at, u.full_name, array_agg(i.image_url) as image_urls
	          FROM violations v
	          JOIN users u ON v.user_id = u.id
	          LEFT JOIN violation_images i ON v.id = i.violation_id
	          GROUP BY v.id, u.id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying all violations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v model.Violation
		var imageUrls []sql.NullString
		if err := rows.Scan(&v.ID, &v.UserID, &v.CarNumber, &v.Description, &v.Status, &v.CreatedAt, &v.FullName, pq.Array(&imageUrls)); err != nil {
			return nil, fmt.Errorf("error scanning violation: %w", err)
		}
		for _, imgUrl := range imageUrls {
			if imgUrl.Valid {
				v.ImageURLs = append(v.ImageURLs, imgUrl.String)
			}
		}
		violations = append(violations, v)
	}

	return violations, nil
}

func (repo *violationRepo) Create(violation model.Violation) (int64, error) {
	// Сначала добавляем запись о нарушении
	query := `INSERT INTO violations (user_id, car_number, description) VALUES ($1, $2, $3) RETURNING id`
	var violationID int64
	err := repo.db.QueryRow(query, violation.UserID, violation.CarNumber, violation.Description).Scan(&violationID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при добавлении заявления о нарушении: %w", err)
	}

	// Теперь добавляем изображения, если они есть
	if len(violation.ImageURLs) > 0 {
		for _, url := range violation.ImageURLs {
			imgQuery := `INSERT INTO violation_images (violation_id, image_url) VALUES ($1, $2)`
			_, imgErr := repo.db.Exec(imgQuery, violationID, url)
			if imgErr != nil {
				// Логирование ошибки без прерывания основной транзакции
				fmt.Errorf("ошибка при добавлении изображения: %w", imgErr)
			}
		}
	}

	return violationID, nil
}

func (repo *violationRepo) FindByUserID(userID int64) ([]model.Violation, error) {
	var violations []model.Violation

	query := `SELECT v.id, v.user_id, v.car_number, v.description, v.status, v.created_at, array_agg(i.image_url) as image_urls
	          FROM violations v
	          LEFT JOIN violation_images i ON v.id = i.violation_id
	          WHERE v.user_id = $1
	          GROUP BY v.id`
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе заявлений: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v model.Violation
		var imageUrls []sql.NullString
		if err := rows.Scan(&v.ID, &v.UserID, &v.CarNumber, &v.Description, &v.Status, &v.CreatedAt, pq.Array(&imageUrls)); err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки заявлений: %w", err)
		}
		for _, imgUrl := range imageUrls {
			if imgUrl.Valid {
				v.ImageURLs = append(v.ImageURLs, imgUrl.String)
			}
		}
		violations = append(violations, v)
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
