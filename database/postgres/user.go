package postgres

import (
	"context"

	"github.com/afe0c1cd/db8c1186/model"
)

func (r *PostgresRepository) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Roles").Preload("Organization").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresRepository) FindAllUser(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.WithContext(ctx).Preload("Roles").Preload("Organization").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return r.FindUserByID(ctx, user.ID.String())
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, err
	}
	return r.FindUserByID(ctx, user.ID.String())
}
