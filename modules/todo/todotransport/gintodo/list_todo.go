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

func ListTodo(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter todomodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInternal(err))
		}

		paging.Fulfill()

		store := todostorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := todobiz.NewListTodoBiz(store)

		result, err := biz.ListTodo(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
