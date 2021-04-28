package todobiz

import (
	"context"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

type CreateTodoStore interface {
	Create(ctx context.Context, data *todomodel.TodoCreate) error
}

type createTodoBiz struct {
	store CreateTodoStore
}

func NewCreateTodoBiz(store CreateTodoStore) *createTodoBiz {
	return &createTodoBiz{store: store}
}

func (biz *createTodoBiz) CreateTodo(ctx context.Context, data *todomodel.TodoCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity(todomodel.EntityName, err)
	}
	return err
}
