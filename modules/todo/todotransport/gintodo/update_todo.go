package gintodo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/common"
	"todo-app/component"
	"todo-app/modules/todo/todobiz"
	"todo-app/modules/todo/todomodel"
	"todo-app/modules/todo/todostorage"
)

func UpdateTodo(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data todomodel.TodoUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := todostorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := todobiz.NewUpdateTodoBiz(store)

		if err := biz.UpdateTodo(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
