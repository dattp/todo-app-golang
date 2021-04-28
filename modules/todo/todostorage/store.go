package todostorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

// NewSQLStore public ra
// can cai gi thi truyen cai do vao
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
