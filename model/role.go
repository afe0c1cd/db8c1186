package model

import (
	"github.com/google/uuid"
)

const (
	RoleNameNone   = "none"
	RoleNameViewer = "viewer"
	RoleNameEditor = "editor"
)

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"size:50;not null;unique" json:"name"`
}
