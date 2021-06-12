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


type JsonRequest struct {
	FieldStr  string `json:"field_str"`
	FieldInt  int    `json:"field_int"`
	FieldBool bool   `json:"field_bool"`
}

func main() {
  db := GetDBConn()

   // テーブルの作成
  db.AutoMigrate(&Assignee{})
  db.AutoMigrate(&Task{})

  r := gin.Default()
  r.GET("/tasks", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
          "message": "tasks",
      })
  })

	r.GET("/asignees", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "asignees",
    })
  })

  r.POST("/asignees", func(c *gin.Context) {
		var json JsonRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"str": json.FieldStr, "int": json.FieldInt, "bool": json.FieldBool})
	})
    
  r.Run(":3001")  // EXPOSE Ports
}