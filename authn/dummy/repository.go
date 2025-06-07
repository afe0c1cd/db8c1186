package dummy

import (
	"github.com/afe0c1cd/db8c1186/authn"
	"github.com/google/uuid"
)

type Repository struct {
}

var (
	OrganizationIDOfA = uuid.MustParse("7f0357c3-6719-4836-a3eb-bcf06ec2d759")
)

var (
	TestTokens = map[string]authn.AuthenticatedUser{
		"a-alice-viewer": { // Alice (A, Inc. - Viewer)
			ID:             uuid.MustParse("fa224131-4ac9-4bc1-ae14-7d5f2c226255"),
			OrganizationID: OrganizationIDOfA,
			Name:           "Alice / Viewer",
		},
		"a-bob-editor": { // Bob (A, Inc. - Editor)
			ID:             uuid.MustParse("6d2dfe34-f2a5-4d34-b865-c2457062cec5"),
			OrganizationID: OrganizationIDOfA,
			Name:           "Bob / Editor",
		},
	}
)

func (r *Repository) AuthenticateByToken(token string) (*authn.AuthenticatedUser, error) {
	u, ok := TestTokens[token]
	if ok {
		return &u, nil
	}
	return nil, authn.ErrInvalidToken
}
