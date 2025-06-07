package postgres

import (
	"context"

	"github.com/afe0c1cd/db8c1186/model"
	"github.com/google/uuid"
)

func (r *PostgresRepository) AddRoleToUser(ctx context.Context, userID, roleID uuid.UUID) error {
	userRole := &model.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	return r.db.WithContext(ctx).Create(userRole).Error
}

func (r *PostgresRepository) RemoveRoleFromUser(ctx context.Context, userID, roleID string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&model.UserRole{}).Error
}

func (r *PostgresRepository) FindRolesByUserID(ctx context.Context, userID string) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).
		Table("roles").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *PostgresRepository) FindUsersByRoleID(ctx context.Context, roleID string) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).
		Table("users").
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Where("user_roles.role_id = ?", roleID).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
