package gintodo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/component"
	"todo-app/modules/todo/todobiz"
	"todo-app/modules/todo/todostorage"
)

func GetTodo(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := todostorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := todobiz.NewGetTodoBi(store)

		data, err := biz.GetTodo(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	}
}
