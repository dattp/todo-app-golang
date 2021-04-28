package todostorage

import (
	"context"
	"todo-app/modules/todo/todomodel"
)

func (s sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *todomodel.TodoUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
