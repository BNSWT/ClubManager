package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TaskApi struct {
}

var taskService = service.ServiceGroupApp.TaskServiceGroup.TaskService

// CreateTask 创建Task
// @Tags Task
// @Summary 创建Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.Task true "创建Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/createTask [post]
func (taskApi *TaskApi) CreateTask(c *gin.Context) {
	var task task.Task
	_ = c.ShouldBindJSON(&task)
	if err := taskService.CreateTask(task); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTask 删除Task
// @Tags Task
// @Summary 删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.Task true "删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTask [delete]
func (taskApi *TaskApi) DeleteTask(c *gin.Context) {
	var task task.Task
	_ = c.ShouldBindJSON(&task)
	if err := taskService.DeleteTask(task); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskByIds 批量删除Task
// @Tags Task
// @Summary 批量删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /task/deleteTaskByIds [delete]
func (taskApi *TaskApi) DeleteTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := taskService.DeleteTaskByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTask 更新Task
// @Tags Task
// @Summary 更新Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.Task true "更新Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateTask [put]
func (taskApi *TaskApi) UpdateTask(c *gin.Context) {
	var task task.Task
	_ = c.ShouldBindJSON(&task)
	if err := taskService.UpdateTask(task); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTask 用id查询Task
// @Tags Task
// @Summary 用id查询Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query task.Task true "用id查询Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTask [get]
func (taskApi *TaskApi) FindTask(c *gin.Context) {
	var task task.Task
	_ = c.ShouldBindQuery(&task)
	if retask, err := taskService.GetTask(task.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retask": retask}, c)
	}
}

// GetTaskList 分页获取Task列表
// @Tags Task
// @Summary 分页获取Task列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query taskReq.TaskSearch true "分页获取Task列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getTaskList [get]
func (taskApi *TaskApi) GetTaskList(c *gin.Context) {
	var pageInfo taskReq.TaskSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := taskService.GetTaskInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
