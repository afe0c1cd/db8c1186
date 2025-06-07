package postgres

import (
	"context"

	"github.com/afe0c1cd/db8c1186/database"
	"github.com/afe0c1cd/db8c1186/model"
	"github.com/google/uuid"
)

func (r *PostgresRepository) FindTodoByID(ctx context.Context, id string) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.WithContext(ctx).Preload("CreatedByUser").Preload("Organization").Preload("Assignees").First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *PostgresRepository) FindAllTodoByUserID(ctx context.Context, userID string, visibility string) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.db.WithContext(ctx).Preload("CreatedByUser").Preload("Organization").Preload("Assignees").Where("created_by_user_id = ?", userID).Where("visibility = ?", visibility).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *PostgresRepository) FindAllTodoByUserIDAndOrganizationID(ctx context.Context, userID string, organizationID string) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.db.WithContext(ctx).Preload("CreatedByUser").Preload("Organization").Preload("Assignees").
		Where(
			"(organization_id = ? AND visibility = 'internal') OR (created_by_user_id = ? AND visibility IN ('private', 'internal'))",
			organizationID, userID,
		).
		Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *PostgresRepository) FindAllTodoByOrganizationID(ctx context.Context, organizationID string, visibility string) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.db.WithContext(ctx).Preload("CreatedByUser").Preload("Organization").Preload("Assignees").Where("organization_id = ?", organizationID).Where("visibility = ?", visibility).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *PostgresRepository) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	if err := r.db.WithContext(ctx).Create(todo).Error; err != nil {
		return nil, err
	}
	return r.FindTodoByID(ctx, todo.ID.String())
}

func (r *PostgresRepository) UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	if err := r.db.WithContext(ctx).Save(todo).Error; err != nil {
		return nil, err
	}
	return r.FindTodoByID(ctx, todo.ID.String())
}

func (r *PostgresRepository) DeleteTodo(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Todo{}, "id = ?", id).Error
}

func (r *PostgresRepository) AssignTodoToUser(ctx context.Context, todoID, userID string) error {
	todo, err := r.FindTodoByID(ctx, todoID)
	if err != nil {
		return err
	}
	if todo == nil {
		return database.ErrNotFound
	}

	todoAssignee := &model.TodoAssignee{
		TodoID: todo.ID,
		UserID: uuid.MustParse(userID),
	}

	if err := r.db.WithContext(ctx).FirstOrCreate(todoAssignee).Error; err != nil {
		return err
	}

	return nil
}
