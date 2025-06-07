package postgres

import (
	"github.com/afe0c1cd/db8c1186/database"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) database.Repository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) WithTx(fn func(tx database.Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := &PostgresRepository{db: tx}
		return fn(txRepo)
	})
}
