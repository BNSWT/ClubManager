// 自动生成模板Task
package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Task 结构体
// 如果含有time.Time 请自行import time包
type Task struct {
	global.GVA_MODEL
	TaskDescription string     `json:"taskDescription" form:"taskDescription" gorm:"column:task_description;comment:;"`
	TaskAssignee    *int       `json:"taskAssignee" form:"taskAssignee" gorm:"column:task_assignee;comment:;"`
	TaskDeadline    *time.Time `json:"taskDeadline" form:"taskDeadline" gorm:"column:task_deadline;comment:;"`
	TaskReport      string     `json:"taskReport" form:"taskReport" gorm:"column:task_report;comment:任务完成后的材料上传，存储文件的名称;"`
	TaskFinished    *bool      `json:"taskFinished" form:"taskFinished" gorm:"column:task_finished;comment:;"`
	ParentTask      *int       `json:"parentTask" form:"parentTask" gorm:"column:parent_task;comment:;"`
	PreTask         *int       `json:"preTask" form:"preTask" gorm:"column:pre_task;comment:;"`
}

// TableName Task 表名
func (Task) TableName() string {
	return "task"
}
