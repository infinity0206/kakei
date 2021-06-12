package main

import "github.com/gin-gonic/gin"

import (
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "net/http"
  // "fmt"
)

func GetDBConn() *gorm.DB {
  db, err := gorm.Open(GetDBConfig())
  if err != nil {
      panic(err)
  }

  db.LogMode(true)
  return db
}

func GetDBConfig() (string, string) {
  DBMS := "mysql"
  USER := "admin"
  PASS := "admin"
  PROTOCOL := "tcp(mysql)"
  DBNAME := "root"
  OPTION := "charset=utf8&parseTime=True&loc=Local"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

  return DBMS, CONNECT
}

type Model struct {
  ID uint `gorm:"primary_key" json:"id"`
  CreatedAt time.Time
	UpdatedAt time.Time
  DeletedAt *time.Time `gorm:"index" json:"-"`
}

type Assignee struct {
	gorm.Model
	Name string `gorm:"not_null;" binding:"required" json:"name"`
  Tasks []*Task `gorm:"many2many:assignee_tasks;" json:"tasks"`
}

type Task struct {
	gorm.Model
	Name string `gorm:"not_null;" binding:"required" json:"name"`
  Assignees []*Assignee `gorm:"many2many:assignee_tasks;" json:"assignees"`
}

func main() {
  db := GetDBConn()

  db.DropTable("assignee_tasks")
  db.DropTable("assignees")
  db.DropTable("tasks")

  // テーブルの作成
  db.AutoMigrate(&Assignee{})
  db.AutoMigrate(&Task{})

  r := gin.Default()
  r.GET("/tasks", func(c *gin.Context) {
    var task []Task
    result := db.Find(&task)
    if result.Error != nil {
      // ここでエラーハンドリング
      c.JSON(http.StatusBadRequest, gin.H{
        "error": result.Error,
      })
			return
    }

    c.JSON(http.StatusOK, gin.H{
      "message": "task",
      "data": result.Value,
    })
  })

  r.POST("/tasks", func(c *gin.Context) {
		var task Task
    // var assignee Assignee
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
        "error": err.Error(),
      })
			return
		}

    // insert処理
    // db.NewRecord(&task);
    db.Create(&task);
    // if err := db.Model(&assignee).Association("tasks").Append(&task).Error; err != nil {
    //   c.JSON(http.StatusBadRequest, gin.H{
    //     "error": err.Error(),
    //   })
    //   return
    // }

		c.JSON(http.StatusOK, gin.H{
      "code": http.StatusOK,
      "status": "ok",
    })
	})

	r.GET("/assignees", func(c *gin.Context) {
    var assignee []Assignee
    result := db.Find(&assignee)
    if result.Error != nil {
      // ここでエラーハンドリング
      c.JSON(http.StatusBadRequest, gin.H{
        "error": result.Error,
      })
			return
    }

    c.JSON(http.StatusOK, gin.H{
      "message": "assignees",
      "data": result,
    })
  })

  r.POST("/assignees", func(c *gin.Context) {
    // var task Task
		var assignee Assignee
		if err := c.ShouldBindJSON(&assignee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
        "error": err.Error(),
      })
			return
		}

    // insert処理
    // db.NewRecord(&assignee);
    db.Create(&assignee);
    // if err := db.Model(&task).Association("assignees").Append(&assignee).Error; err != nil {
    //   c.JSON(http.StatusBadRequest, gin.H{
    //     "error": err.Error(),
    //   })
    //   return
    // }

		c.JSON(http.StatusOK, gin.H{
      "code": http.StatusOK,
      "status": "ok",
    })
	})
    
  r.Run(":3001")  // EXPOSE Ports
}