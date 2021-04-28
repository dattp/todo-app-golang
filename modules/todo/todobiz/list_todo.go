package todobiz

import (
	"context"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

type ListTodoStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		filter *todomodel.Filter, paging *common.Paging, moreKey ...string) ([]todomodel.Todo, error)
}

type listTodoBiz struct {
	store ListTodoStore
}

func NewListTodoBiz(store ListTodoStore) *listTodoBiz {
	return &listTodoBiz{store: store}
}

func (biz *listTodoBiz) ListTodo(ctx context.Context,
	filter *todomodel.Filter,
	paging *common.Paging,
) ([]todomodel.Todo, error) {

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(todomodel.EntityName, err)
	}

	return result, err
}
