package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
	"todo-app/component"
	"todo-app/middleware"
	"todo-app/modules/todo/todotransport/gintodo"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db, err)
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {

	appCxt := component.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCxt))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	todo := r.Group("/todos")
	{
		todo.POST("", gintodo.CreateTodo(appCxt))
		todo.GET("", gintodo.ListTodo(appCxt))
		todo.GET("/:id", gintodo.GetTodo(appCxt))
		todo.PATCH("/:id", gintodo.UpdateTodo(appCxt))
		todo.DELETE("/:id", gintodo.DeleteTodo(appCxt))

	}

	return r.Run("localhost:5000")
}
