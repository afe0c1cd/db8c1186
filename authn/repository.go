package authn

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type AuthenticatedUser struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
}

type Repository interface {
	AuthenticateByToken(token string) (*AuthenticatedUser, error)
}
