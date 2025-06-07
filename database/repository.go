package database

import (
	"context"
	"errors"

	"github.com/afe0c1cd/db8c1186/model"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	UserRepository
	OrganizationRepository
	RoleRepository
	TodoRepository

	WithTx(fn func(tx Repository) error) error
}

type UserRepository interface {
	FindUserByID(ctx context.Context, id string) (*model.User, error)
	FindAllUser(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type OrganizationRepository interface {
	FindOrganizationByID(ctx context.Context, id string) (*model.Organization, error)
	FindAllOrganization(ctx context.Context) ([]*model.Organization, error)
	CreateOrganization(ctx context.Context, organization *model.Organization) (*model.Organization, error)
	UpdateOrganization(ctx context.Context, organization *model.Organization) (*model.Organization, error)
}

type RoleRepository interface {
	FindRoleByID(ctx context.Context, id string) (*model.Role, error)
	FindAllRole(ctx context.Context) ([]*model.Role, error)
	CreateRole(ctx context.Context, role *model.Role) (*model.Role, error)
	UpdateRole(ctx context.Context, role *model.Role) (*model.Role, error)
}

type TodoRepository interface {
	FindTodoByID(ctx context.Context, id string) (*model.Todo, error)
	FindAllTodoByUserID(ctx context.Context, userID string, visibility string) ([]*model.Todo, error)
	FindAllTodoByUserIDAndOrganizationID(ctx context.Context, userID string, organizationID string) ([]*model.Todo, error)
	FindAllTodoByOrganizationID(ctx context.Context, organizationID string, visibility string) ([]*model.Todo, error)
	CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	DeleteTodo(ctx context.Context, id string) error

	AssignTodoToUser(ctx context.Context, todoID, userID string) error
}

type UserRoleRepository interface {
	AddRoleToUser(ctx context.Context, userID, roleID string) error
	RemoveRoleFromUser(ctx context.Context, userID, roleID string) error
	FindRolesByUserID(ctx context.Context, userID string) ([]*model.Role, error)
	FindUsersByRoleID(ctx context.Context, roleID string) ([]*model.User, error)
}
