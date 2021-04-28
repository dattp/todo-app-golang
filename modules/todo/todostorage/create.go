package todostorage

import (
	"context"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

func (s *sqlStore) Create(ctx context.Context, data *todomodel.TodoCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
