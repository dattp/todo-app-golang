package todostorage

import (
	"context"
	"todo-app/modules/todo/todomodel"
)

func (s sqlStore) SoftDeleteData(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(todomodel.Todo{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return err
	}
	return nil
}
