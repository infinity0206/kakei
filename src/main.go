package main

import (
	"main/db"
	"main/server"
)

func main() {
	db.Init()
	server.Init()
}

// package main

// import "github.com/gin-gonic/gin"

// import (
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"main/db"
// 	"main/entity"
// 	"net/http"
//   "main/controller"
// )

// type Assignee entity.Assignee
// type Task entity.Task
// type AssigneeTask entity.AssigneeTask

// func main() {
// 	db.Init()
// 	db := db.GetDB()

// 	// db.DropTable("assignee_tasks")
// 	// db.DropTable("assignees")
// 	// db.DropTable("tasks")

// 	// // テーブルの作成
// 	// db.AutoMigrate(&Assignee{})
// 	// db.AutoMigrate(&Task{})
// 	// db.AutoMigrate(&AssigneeTask{})

// 	// r := gin.Default()
// 	// r.GET("/tasks", func(c *gin.Context) {
// 	// 	var task []Task
// 	// 	result := db.Preload("Assignees").Find(&task)
// 	// 	if result.Error != nil {
// 	// 		c.JSON(http.StatusBadRequest, gin.H{
// 	// 			"error": result.Error,
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"message": "task",
// 	// 		"data":    result.Value,
// 	// 	})
// 	// })

// 	// r.POST("/tasks", func(c *gin.Context) {
// 	// 	var task Task
// 	// 	if err := c.ShouldBindJSON(&task); err != nil {
// 	// 		c.JSON(http.StatusBadRequest, gin.H{
// 	// 			"error": err.Error(),
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	db.Create(&task)

// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"code":   http.StatusOK,
// 	// 		"status": "ok",
// 	// 	})
// 	// })

// 	// r.GET("/assignees", func(c *gin.Context) {
// 	// 	var assignee []Assignee
// 	// 	result := db.Preload("Tasks").Find(&assignee)
// 	// 	if result.Error != nil {
// 	// 		c.JSON(http.StatusBadRequest, gin.H{
// 	// 			"error": result.Error,
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"message": "assignees",
// 	// 		"data":    result.Value,
// 	// 	})
// 	// })

// 	// r.POST("/assignees", func(c *gin.Context) {
// 	// 	var assignee Assignee
// 	// 	if err := c.ShouldBindJSON(&assignee); err != nil {
// 	// 		c.JSON(http.StatusBadRequest, gin.H{
// 	// 			"error": err.Error(),
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	db.Create(&assignee)

// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"code":   http.StatusOK,
// 	// 		"status": "ok",
// 	// 	})
// 	// })

// 	r.Run(":3001") // EXPOSE Ports
// }
