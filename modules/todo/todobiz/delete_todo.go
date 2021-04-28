package todobiz

import (
	"context"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

type DeleteTodoStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKey ...string,
	) (*todomodel.Todo, error)

	SoftDeleteData(ctx context.Context, id int) error
}

type deleteTodoBiz struct {
	store DeleteTodoStore
}

func NewDeleteTodoBiz(store DeleteTodoStore) *deleteTodoBiz {
	return &deleteTodoBiz{store: store}
}

func (biz *deleteTodoBiz) DeleteTodo(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(todomodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(todomodel.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(todomodel.EntityName, err)
	}
	return nil
}
