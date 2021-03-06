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

func DeleteTodo(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := todostorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := todobiz.NewDeleteTodoBiz(store)

		if err := biz.DeleteTodo(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
