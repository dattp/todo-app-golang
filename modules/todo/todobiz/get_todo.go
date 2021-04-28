package todobiz

import (
	"context"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

type GetTodoStore interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*todomodel.Todo, error)
}

type getTodoBiz struct {
	store GetTodoStore
}

func NewGetTodoBi(store GetTodoStore) *getTodoBiz {
	return &getTodoBiz{store: store}
}

func (biz *getTodoBiz) GetTodo(ctx context.Context, id int) (*todomodel.Todo, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {

		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(todomodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(todomodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(todomodel.EntityName, nil)
	}

	return data, err
}
