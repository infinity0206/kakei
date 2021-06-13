package db

import (
  "os"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/joho/godotenv"
  "main/entity"
  // "fmt"
)

var (
  db *gorm.DB
  err error
)

type Assignee entity.Assignee
type Task entity.Task
type AssigneeTask entity.AssigneeTask

// DB初期化
func Init() {
  // 環境変数取得
  godotenv.Load(".env")
  
  CONNECT := os.Getenv("USER") + ":" + os.Getenv("PASSWORD") + "@tcp(" + os.Getenv("HOST") + ")/" + os.Getenv("DATABASE")  + "?" + os.Getenv("OPTION")
  // DB接続
  db, err = gorm.Open("mysql", CONNECT)

  if err != nil {
    panic(err)
  }

  DropTables()
  autoMigration()
}

// func GetDBConn() *gorm.DB {
// 	db, err := gorm.Open(GetDBConfig())
// 	if err != nil {
// 		panic(err)
// 	}
  
// 	db.LogMode(true)
// 	return db
//   }
  
// func GetDBConfig() (string, string) {
// 	DBMS := "mysql"
// 	USER := "admin"
// 	PASS := "admin"
// 	PROTOCOL := "tcp(mysql)"
// 	DBNAME := "root"
// 	OPTION := "charset=utf8&parseTime=True&loc=Local"
  
// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
  
// 	return DBMS, CONNECT
// }

// DB取得
func GetDB() *gorm.DB {
  return db
}

// DB接続終了
func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}

// テーブルのドロップ
func DropTables() {
  db.DropTable("assignee_tasks")
  db.DropTable("assignees")
  db.DropTable("tasks")
}

// マイグレーション
func autoMigration() {
  db.AutoMigrate(&Assignee{})
  db.AutoMigrate(&Task{})
  db.AutoMigrate(&AssigneeTask{})
}

