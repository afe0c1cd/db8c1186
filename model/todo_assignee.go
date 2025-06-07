package model

import (
	"time"

	"github.com/google/uuid"
)

type TodoAssignee struct {
	UserID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	TodoID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"todo_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedByUserID uuid.UUID `gorm:"type:uuid;not null" json:"created_by_user_id"`

	Todo Todo `gorm:"foreignKey:TodoID;" json:"todo,omitempty"`
	User User `gorm:"foreignKey:UserID;" json:"user,omitempty"`
}
