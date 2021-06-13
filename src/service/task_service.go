package service

import (
	"main/db"
	"main/entity"
)

type TaskService struct{}

// 全取得
func (s TaskService) FindAll() ([]entity.Task, error) {

	// DB接続
	db := db.GetDB()

	// 本モデルから作成
	var task []entity.Task

	// DB接続確認
	// if err := db.Take(&task).Error; err != nil {
	// 	return task, err
	// }

	db.Preload("Assignees").Find(&task)

	return task, nil
}

// 登録
func (s TaskService) Create() (entity.Task, error) {
	// DB接続
	db := db.GetDB()

	// 本モデルから作成
	var task entity.Task
	var assignee entity.Assignee

	// DB接続確認
	// if err := db.Take(&task).Error; err != nil {
	// 	return task, err
	// }

	db.Create(&task)
	db.Model(&task).Association("Assignees").Append(&assignee)

	return task, nil
}