package repository

import "exambackend/internal/model"

// ViolationRepository определяет интерфейс для работы со заявлениями о нарушениях
type ViolationRepository interface {
	Create(violation model.Violation) (int64, error)
	FindByUserID(userID int64) ([]model.Violation, error)
	UpdateStatus(violationID int64, status string) error
	GetAllViolations() ([]model.Violation, error)
}
