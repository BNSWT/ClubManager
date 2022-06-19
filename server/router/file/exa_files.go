package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExaFilesRouter struct {
}

// InitExaFilesRouter 初始化 ExaFiles 路由信息
func (s *ExaFilesRouter) InitExaFilesRouter(Router *gin.RouterGroup) {
	exaFilesRouter := Router.Group("exaFiles").Use(middleware.OperationRecord())
	exaFilesRouterWithoutRecord := Router.Group("exaFiles")
	var exaFilesApi = v1.ApiGroupApp.FileApiGroup.ExaFilesApi
	{
		exaFilesRouter.POST("createExaFiles", exaFilesApi.CreateExaFiles)   // 新建ExaFiles
		exaFilesRouter.DELETE("deleteExaFiles", exaFilesApi.DeleteExaFiles) // 删除ExaFiles
		exaFilesRouter.DELETE("deleteExaFilesByIds", exaFilesApi.DeleteExaFilesByIds) // 批量删除ExaFiles
		exaFilesRouter.PUT("updateExaFiles", exaFilesApi.UpdateExaFiles)    // 更新ExaFiles
	}
	{
		exaFilesRouterWithoutRecord.GET("findExaFiles", exaFilesApi.FindExaFiles)        // 根据ID获取ExaFiles
		exaFilesRouterWithoutRecord.GET("getExaFilesList", exaFilesApi.GetExaFilesList)  // 获取ExaFiles列表
	}
}
