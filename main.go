package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"todo-app/component"
	"todo-app/component/uploadprovider"
	"todo-app/middleware"
	"todo-app/modules/todo/todotransport/gintodo"
	"todo-app/modules/upload/uploadtransport/ginupload"
)

func main() {

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	dsn := os.Getenv("DBConnectionStr")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db, err)
	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) error {

	appCxt := component.NewAppContext(db, uploadProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCxt))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", ginupload.Upload(appCxt))

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
