package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRouter struct {
}

// InitTaskRouter 初始化 Task 路由信息
func (s *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("task").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("task")
	var taskApi = v1.ApiGroupApp.TaskApiGroup.TaskApi
	{
		taskRouter.POST("createTask", taskApi.CreateTask)   // 新建Task
		taskRouter.DELETE("deleteTask", taskApi.DeleteTask) // 删除Task
		taskRouter.DELETE("deleteTaskByIds", taskApi.DeleteTaskByIds) // 批量删除Task
		taskRouter.PUT("updateTask", taskApi.UpdateTask)    // 更新Task
	}
	{
		taskRouterWithoutRecord.GET("findTask", taskApi.FindTask)        // 根据ID获取Task
		taskRouterWithoutRecord.GET("getTaskList", taskApi.GetTaskList)  // 获取Task列表
	}
}
