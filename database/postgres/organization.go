package postgres

import (
	"context"

	"github.com/afe0c1cd/db8c1186/model"
)

func (r *PostgresRepository) FindOrganizationByID(ctx context.Context, id string) (*model.Organization, error) {
	var org model.Organization
	if err := r.db.WithContext(ctx).First(&org, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (r *PostgresRepository) FindAllOrganization(ctx context.Context) ([]*model.Organization, error) {
	var orgs []*model.Organization
	if err := r.db.WithContext(ctx).Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}

func (r *PostgresRepository) CreateOrganization(ctx context.Context, org *model.Organization) (*model.Organization, error) {
	if err := r.db.WithContext(ctx).Create(org).Error; err != nil {
		return nil, err
	}
	return org, nil
}

func (r *PostgresRepository) UpdateOrganization(ctx context.Context, org *model.Organization) (*model.Organization, error) {
	if err := r.db.WithContext(ctx).Save(org).Error; err != nil {
		return nil, err
	}
	return org, nil
}

func (r *PostgresRepository) DeleteOrganization(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Organization{}, "id = ?", id).Error
}
