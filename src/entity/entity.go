package entity

import (
	"github.com/jinzhu/gorm"
)

type Assignee struct {
	gorm.Model
	Name  string  `gorm:"not_null;" binding:"required" json:"name"`
	Tasks []*Task `gorm:"many2many:assignee_tasks;" json:"tasks"`
}

type Task struct {
	gorm.Model
	Name      string      `gorm:"not_null;" binding:"required" json:"name"`
	Assignees []*Assignee `gorm:"many2many:assignee_tasks;" json:"assignees"`
}

type AssigneeTask struct {
	AssigneeId int `gorm:"not_null;" binding:"required" json:"assignee_id"`
	TaskId     int `gorm:"not_null;" binding:"required" json:"task_id"`
}