package main

import (
	"github.com/firstnapat/todo/auth"
	"github.com/firstnapat/todo/todo"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tokenz", auth.AccessToken)

	handler := todo.NewTodoHandler(db)
	r.POST("/todos", handler.NewTask)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
