package todobiz

import (
	"context"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

type UpdateTodoStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKey ...string,
	) (*todomodel.Todo, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *todomodel.TodoUpdate,
	) error
}

type updateTodoBiz struct {
	store UpdateTodoStore
}

func NewUpdateTodoBiz(store UpdateTodoStore) *updateTodoBiz {
	return &updateTodoBiz{store: store}
}

func (biz *updateTodoBiz) UpdateTodo(ctx context.Context, id int, data *todomodel.TodoUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(todomodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(todomodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(todomodel.EntityName, err)
	}

	return nil
}
