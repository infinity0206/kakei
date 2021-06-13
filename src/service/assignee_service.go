package service

import (
	"main/db"
	"main/entity"
	// "fmt"
)

type AssigneeService struct{}


// 全取得
func (s AssigneeService) FindAll() ([]entity.Assignee, error) {

	// DB接続
	db := db.GetDB()

	// 本モデルから作成
	var assignee []entity.Assignee

	// DB接続確認
	// if err := db.Take(&assignee).Error; err != nil {
	// 	return nil, err
	// }
	db.Preload("Tasks").Find(&assignee)

	return assignee, nil
}

// 登録
func (s AssigneeService) Create() (entity.Assignee, error) {
	// DB接続
	db := db.GetDB()

	// 本モデルから作成
	var assignee entity.Assignee
	var task entity.Task

	// DB接続確認
	// if err := db.Take(&assignee).Error; err != nil {
	// 	return assignee, err
	// }

	db.Create(&assignee)
	db.Model(&assignee).Association("Tasks").Append(&task)

	return assignee, nil
}