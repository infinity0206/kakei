package entity

type AssigneeTask struct {
	AssigneeId int `gorm:"not_null;" binding:"required" json:"assignee_id"`
	TaskId     int `gorm:"not_null;" binding:"required" json:"task_id"`
}
