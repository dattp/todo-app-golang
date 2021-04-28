package gintodo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/common"
	"todo-app/component"
	"todo-app/modules/todo/todobiz"
	"todo-app/modules/todo/todomodel"
	"todo-app/modules/todo/todostorage"
)

func CreateTodo(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data todomodel.TodoCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := todostorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := todobiz.NewCreateTodoBiz(store)

		if err := biz.CreateTodo(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
