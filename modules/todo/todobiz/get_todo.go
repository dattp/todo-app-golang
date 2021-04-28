package todobiz

import (
	"context"
	"errors"
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
			return nil, err
		}
		return nil, err
	}

	if data.Status == 0 {
		return nil, errors.New("todo deleted")
	}

	return data, err
}
