package model

import "github.com/google/uuid"

type Organization struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"size:255;not null" json:"name"`

	Users []User `gorm:"many2many:organization_users;"`
}
