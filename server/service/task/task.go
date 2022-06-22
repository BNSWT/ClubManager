package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
)

type TaskService struct {
}

// CreateTask 创建Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) CreateTask(task task.Task) (err error) {
	err = global.GVA_DB.Create(&task).Error
	return err
}

// DeleteTask 删除Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)DeleteTask(task task.Task) (err error) {
	err = global.GVA_DB.Delete(&task).Error
	return err
}

// DeleteTaskByIds 批量删除Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)DeleteTaskByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]task.Task{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTask 更新Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)UpdateTask(task task.Task) (err error) {
	err = global.GVA_DB.Save(&task).Error
	return err
}

// GetTask 根据id获取Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)GetTask(id uint) (task task.Task, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&task).Error
	return
}

// GetTaskInfoList 分页获取Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)GetTaskInfoList(info taskReq.TaskSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&task.Task{})
    var tasks []task.Task
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&tasks).Error
	return  tasks, total, err
}
