package server

import (
	"github.com/afe0c1cd/db8c1186/generated"
	"github.com/afe0c1cd/db8c1186/model"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func ToTodo(todo *model.Todo) *generated.Todo {
	assigneeUserIds := make([]openapi_types.UUID, len(todo.Assignees))

	for i, assignee := range todo.Assignees {
		assigneeUserIds[i] = assignee.ID
	}

	var status generated.TodoStatus
	if todo.Status != nil {
		status = generated.TodoStatus(*todo.Status)
	}

	t := &generated.Todo{
		AssigneeUserIds: &assigneeUserIds,
		CreatedAt:       todo.CreatedAt,
		Description:     todo.Description,
		DueDate:         todo.DueDate,
		Id:              todo.ID,
		Status:          &status,
		Title:           todo.Title,
		UpdatedAt:       todo.UpdatedAt,
		Visibility:      generated.TodoVisibility(todo.Visibility),
	}

	return t
}
