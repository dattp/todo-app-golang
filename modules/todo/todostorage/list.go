package todostorage

import (
	"context"
	"todo-app/common"
	"todo-app/modules/todo/todomodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *todomodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]todomodel.Todo, error) {
	var result []todomodel.Todo
	db := s.db
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	db = db.Table(todomodel.Todo{}.TableName()).Where(condition).Where("status in (1)")

	if v := filter; v != nil {
		if v.Status > 0 {
			db = db.Where("status=?", v.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
