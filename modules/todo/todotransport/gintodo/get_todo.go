package gintodo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/common"
	"todo-app/component"
	"todo-app/modules/todo/todobiz"
	"todo-app/modules/todo/todostorage"
)

func GetTodo(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := todostorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := todobiz.NewGetTodoBi(store)

		data, err := biz.GetTodo(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}
