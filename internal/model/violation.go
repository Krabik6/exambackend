package model

import "time"

// Violation представляет заявление о нарушении
type Violation struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	CarNumber   string    `json:"carNumber"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	FullName    string    `json:"fullName"` // Добавлено новое поле
}
