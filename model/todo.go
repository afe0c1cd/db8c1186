package model

import (
	"time"

	"github.com/google/uuid"
)

var (
	TodoVisibilityPrivate  = "private"
	TodoVisibilityInternal = "internal"
)

type Todo struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DueDate         *time.Time `json:"due_date,omitempty"`
	Title           string     `gorm:"size:255;not null" json:"title"`
	Description     *string    `gorm:"type:text" json:"description,omitempty"`
	Status          *string    `gorm:"size:50" json:"status"`
	Visibility      string     `gorm:"size:50;not null" json:"visibility"`
	CreatedByUserID uuid.UUID  `gorm:"type:uuid;not null" json:"created_by_user_id"`
	OrganizationID  uuid.UUID  `gorm:"type:uuid;not null" json:"organization_id"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	CreatedByUser User         `gorm:"foreignKey:CreatedByUserID;" json:"created_by_user,omitempty"`
	Organization  Organization `gorm:"foreignKey:OrganizationID;" json:"organization,omitempty"`
	Assignees     []User       `gorm:"many2many:todo_assignees"`
}
