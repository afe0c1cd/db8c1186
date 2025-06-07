package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name           string    `gorm:"size:255;not null" json:"name"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null" json:"organization_id"`

	Organization *Organization `gorm:"foreignKey:OrganizationID" json:"organization"`
	Todos        []Todo        `gorm:"foreignKey:CreatedByUserID;references:ID" json:"todos,omitempty"`
	Roles        []Role        `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}
