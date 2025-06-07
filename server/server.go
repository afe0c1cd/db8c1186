package server

import (
	"context"
	"net/http"

	"github.com/afe0c1cd/db8c1186/authn/dummy"
	"github.com/afe0c1cd/db8c1186/database"
	"github.com/afe0c1cd/db8c1186/generated"
	md "github.com/afe0c1cd/db8c1186/middleware"
	"github.com/afe0c1cd/db8c1186/model"
	"github.com/afe0c1cd/db8c1186/server/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Server struct {
	e *echo.Echo

	repository database.Repository
}

func NewServer(repo database.Repository) *Server {
	return &Server{
		repository: repo,
	}
}

func (s *Server) Start(addr string) error {
	s.e = echo.New()

	s.e.HidePort = true
	s.e.HideBanner = true

	s.e.Use(middleware.Recover())
	s.e.Use(middleware.Logger())
	s.e.Use(md.AuthenticationMiddleware(s.repository, &dummy.Repository{}))

	s.e.HTTPErrorHandler = md.CustomErrorHandler

	generated.RegisterHandlers(s.e, s)

	return s.e.Start(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s *Server) GetV1Todos(ctx echo.Context, params generated.GetV1TodosParams) error {
	u := GetUser(ctx)

	v := generated.GetV1TodosParamsVisibilityAll
	if params.Visibility != nil {
		v = *params.Visibility
	}

	var todos []*generated.Todo
	var t []*model.Todo
	var err error

	switch v {
	case generated.GetV1TodosParamsVisibilityAll:
		t, err = s.repository.FindAllTodoByUserIDAndOrganizationID(ctx.Request().Context(), u.ID.String(), u.OrganizationID.String())
	case generated.GetV1TodosParamsVisibilityInternal:
		t, err = s.repository.FindAllTodoByOrganizationID(ctx.Request().Context(), u.OrganizationID.String(), model.TodoVisibilityInternal)
	case generated.GetV1TodosParamsVisibilityPrivate:
		t, err = s.repository.FindAllTodoByUserID(ctx.Request().Context(), u.ID.String(), model.TodoVisibilityPrivate)
	}
	if err != nil {
		return errors.NewInternalServerError(err)
	}

	for _, todo := range t {
		todos = append(todos, ToTodo(todo))
	}

	return ctx.JSON(http.StatusOK, todos)
}

func (s *Server) PostV1Todos(ctx echo.Context) error {
	var todo generated.CreateTodoRequest
	if err := ctx.Bind(&todo); err != nil {
		return errors.NewBadRequest("invalid request body")
	}

	if !CanAddTodo(ctx) {
		return errors.NewForbidden("user does not have permission to create todos")
	}

	var t *model.Todo
	err := s.repository.WithTx(
		func(tx database.Repository) error {
			var err error
			t, err = tx.CreateTodo(ctx.Request().Context(), &model.Todo{
				Title:           todo.Title,
				Description:     todo.Description,
				DueDate:         todo.DueDate,
				Visibility:      string(todo.Visibility),
				Status:          (*string)(todo.Status),
				CreatedByUserID: GetUser(ctx).ID,
				OrganizationID:  GetUser(ctx).OrganizationID,
			})
			if err != nil {
				return err
			}

			if todo.AssigneeUserIds != nil && len(*todo.AssigneeUserIds) > 0 {
				for _, assigneeID := range *todo.AssigneeUserIds {
					if err := tx.AssignTodoToUser(ctx.Request().Context(), t.ID.String(), assigneeID.String()); err != nil {
						return err
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return errors.NewInternalServerError(err)
	}

	return ctx.JSON(http.StatusCreated, ToTodo(t))
}

func (s *Server) DeleteV1TodosTodoId(ctx echo.Context, todoId openapi_types.UUID) error {
	todo, err := s.repository.FindTodoByID(ctx.Request().Context(), todoId.String())
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	if todo == nil {
		return errors.TodoNotFound()
	}

	if !CanEditOrDeleteTodo(ctx, todo) {
		return errors.TodoNotFound()
	}

	err = s.repository.WithTx(
		func(tx database.Repository) error {
			return tx.DeleteTodo(ctx.Request().Context(), todoId.String())
		},
	)
	if err != nil {
		return errors.NewInternalServerError(err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) PatchV1TodosTodoId(ctx echo.Context, todoId openapi_types.UUID) error {
	var todo generated.UpdateTodoRequest
	if err := ctx.Bind(&todo); err != nil {
		return errors.NewBadRequest("invalid request body")
	}

	t, err := s.repository.FindTodoByID(ctx.Request().Context(), todoId.String())
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	if t == nil {
		return errors.TodoNotFound()
	}

	if !CanEditOrDeleteTodo(ctx, t) {
		return errors.TodoNotFound()
	}

	if todo.Title != nil {
		t.Title = *todo.Title
	}
	if todo.Description != nil {
		t.Description = todo.Description
	}
	if todo.DueDate != nil {
		t.DueDate = todo.DueDate
	}
	if todo.Visibility != nil {
		t.Visibility = string(*todo.Visibility)
	}
	if todo.Status != nil {
		status := string(*todo.Status)
		t.Status = &status
	}

	err = s.repository.WithTx(
		func(tx database.Repository) error {
			_, err := tx.UpdateTodo(ctx.Request().Context(), t)
			return err
		},
	)
	if err != nil {
		return errors.NewInternalServerError(err)
	}

	return ctx.JSON(http.StatusOK, ToTodo(t))
}

func CanAddTodo(ctx echo.Context) bool {
	u := GetUser(ctx)
	for _, role := range u.Roles {
		if role.Name == model.RoleNameEditor {
			return true
		}
	}
	return false
}

func CanEditOrDeleteTodo(ctx echo.Context, todo *model.Todo) bool {
	u := GetUser(ctx)
	for _, role := range u.Roles {
		if role.Name == model.RoleNameEditor {
			return true
		}
	}

	if todo.OrganizationID != u.OrganizationID {
		return false
	}

	if todo.Visibility == model.TodoVisibilityInternal {
		return true
	}

	if todo.Visibility == model.TodoVisibilityPrivate && todo.CreatedByUserID == u.ID {
		return true
	}

	return false
}
