package service

import (
	"exambackend/internal/model"
	"exambackend/internal/repository"
	"time"
)

// ViolationService предоставляет методы для работы со заявлениями о нарушениях
type ViolationService interface {
	CreateViolation(violation model.Violation, userId int64) (int64, error)
	GetViolationsByUser(userId int64) ([]model.Violation, error)
	UpdateViolationStatus(violationId int64, status string) error
	GetAllViolations() ([]model.Violation, error)
}

type violationService struct {
	violationRepo repository.ViolationRepository
}

// NewViolationService создает новый экземпляр ViolationService
func NewViolationService(violationRepo repository.ViolationRepository) ViolationService {
	return &violationService{violationRepo: violationRepo}
}

// CreateViolation создает новое заявление о нарушении
func (s *violationService) CreateViolation(violation model.Violation, userId int64) (int64, error) {
	violation.UserID = userId
	violation.CreatedAt = time.Now()
	return s.violationRepo.Create(violation)
}

// GetViolationsByUser возвращает список заявлений, поданных пользователем
func (s *violationService) GetViolationsByUser(userId int64) ([]model.Violation, error) {
	return s.violationRepo.FindByUserID(userId)
}

// UpdateViolationStatus изменяет статус заявления
func (s *violationService) UpdateViolationStatus(violationId int64, status string) error {
	return s.violationRepo.UpdateStatus(violationId, status)
}

func (s *violationService) GetAllViolations() ([]model.Violation, error) {
	return s.violationRepo.GetAllViolations()
}
