package main

import "github.com/gin-gonic/gin"

import (
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "net/http"
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
	Model
	Name string `gorm:"unique;not_null;" binding:"required" json:"name"`
}

type Task struct {
	Model
	Name string `gorm:"unique;not_null;" binding:"required" json:"name"`
	Type int `gorm:"unique;not_null;" binding:"required" json:"type"`
  AssigneeRefer int
  Assignee Assignee `gorm:"foreignKey:AssigneeRefer"`
}

func main() {
  db := GetDBConn()

  db.DropTable("assignees")
  db.DropTable("tasks")
   // テーブルの作成
  db.AutoMigrate(&Assignee{})
  db.AutoMigrate(&Task{})

  r := gin.Default()
  r.GET("/tasks", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "tasks",
    })
  })

	r.GET("/assignees", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "asignees",
    })
  })

  r.POST("/assignees", func(c *gin.Context) {
		var assignee Assignee
		if err := c.ShouldBindJSON(&assignee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
        "error": err.Error(),
      })
			return
		}

    // insert処理
    db.NewRecord(&assignee);
    db.Create(&assignee);
		c.JSON(http.StatusOK, gin.H{
      "code": http.StatusOK,
      "status": "ok",
    })
	})
    
  r.Run(":3001")  // EXPOSE Ports
}