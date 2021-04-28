package todostorage

import (
	"context"
	"gorm.io/gorm"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

func (s *sqlStore) FindDataByCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*todomodel.Todo, error) {
	var result todomodel.Todo
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
