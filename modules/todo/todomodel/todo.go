package todomodel

import (
	"errors"
	"strings"
	"todo-app/common"
)

const EntityName = "Todo"

type Todo struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
	Detail          string `json:"detail" gorm:"column:detail"`
}

func (Todo) TableName() string {
	return "todos"
}

type TodoUpdate struct {
	Title  *string `json:"title" gorm:"column:title"`
	Detail *string `json:"detail" gorm:"column:detail"`
}

func (TodoUpdate) TableName() string {
	return Todo{}.TableName()
}

// TodoCreate vi struct tu csdl lên datamodel có thể nhiều hơn là struct create nên sẽ tạo riêng ra
type TodoCreate struct {
	Id     int    `json:"id" gorm:"column:id"`
	Title  string `json:"title" gorm:"column:title"`
	Detail string `json:"detail" gorm:"column:detail"`
}

func (TodoCreate) TableName() string {
	return Todo{}.TableName()
}

func (todo *TodoCreate) Validate() error {
	todo.Title = strings.TrimSpace(todo.Title)

	if len(todo.Title) == 0 {
		return errors.New("title is empty")
	}

	return nil
}
