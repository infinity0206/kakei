package main

import "github.com/gin-gonic/gin"

import (
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
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
  DeletedAt *time.Time `sql:"index" json:"-"`
}

type Assignee struct {
	Model
	Name string
}

type Task struct {
	Model
	Name string
	Type int
  AssigneeRefer int
  Assignee Assignee `gorm:"foreignKey:AssigneeRefer"`
}

func main() {
  db := GetDBConn()

   // テーブルの作成
  db.AutoMigrate(&Assignee{})
  db.AutoMigrate(&Task{})

  r := gin.Default()
  r.GET("/tasks", func(c *gin.Context) {
      c.JSON(200, gin.H{
          "message": "tasks",
      })
  })

	r.GET("/asignees", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "asignees",
    })
  })
    
  r.Run(":3001")  // EXPOSE Ports
}