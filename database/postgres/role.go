package postgres

import (
	"context"

	"github.com/afe0c1cd/db8c1186/model"
)

func (r *PostgresRepository) FindRoleByID(ctx context.Context, id string) (*model.Role, error) {
	var role model.Role
	if err := r.db.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *PostgresRepository) FindAllRole(ctx context.Context) ([]*model.Role, error) {
	var roles []*model.Role
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *PostgresRepository) CreateRole(ctx context.Context, role *model.Role) (*model.Role, error) {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *PostgresRepository) UpdateRole(ctx context.Context, role *model.Role) (*model.Role, error) {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
