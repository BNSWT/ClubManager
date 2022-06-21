package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {
}

// InitFileRouter 初始化 File 路由信息
func (s *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file").Use(middleware.OperationRecord())
	fileRouterWithoutRecord := Router.Group("file")
	var fileApi = v1.ApiGroupApp.FileApiGroup.FileApi
	{
		fileRouter.POST("createFile", fileApi.CreateFile)   // 新建File
		fileRouter.DELETE("deleteFile", fileApi.DeleteFile) // 删除File
		fileRouter.DELETE("deleteFileByIds", fileApi.DeleteFileByIds) // 批量删除File
		fileRouter.PUT("updateFile", fileApi.UpdateFile)    // 更新File
	}
	{
		fileRouterWithoutRecord.GET("findFile", fileApi.FindFile)        // 根据ID获取File
		fileRouterWithoutRecord.GET("getFileList", fileApi.GetFileList)  // 获取File列表
	}
}
