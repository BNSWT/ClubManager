package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileApi struct {
}

var fileService = service.ServiceGroupApp.FileServiceGroup.FileService

// CreateFile 创建File
// @Tags File
// @Summary 创建File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.File true "创建File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /file/createFile [post]
func (fileApi *FileApi) CreateFile(c *gin.Context) {
	var file file.File
	_ = c.ShouldBindJSON(&file)
	if err := fileService.CreateFile(file); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFile 删除File
// @Tags File
// @Summary 删除File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.File true "删除File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /file/deleteFile [delete]
func (fileApi *FileApi) DeleteFile(c *gin.Context) {
	var file file.File
	_ = c.ShouldBindJSON(&file)
	if err := fileService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileByIds 批量删除File
// @Tags File
// @Summary 批量删除File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /file/deleteFileByIds [delete]
func (fileApi *FileApi) DeleteFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := fileService.DeleteFileByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFile 更新File
// @Tags File
// @Summary 更新File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.File true "更新File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /file/updateFile [put]
func (fileApi *FileApi) UpdateFile(c *gin.Context) {
	var file file.File
	_ = c.ShouldBindJSON(&file)
	if err := fileService.UpdateFile(file); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFile 用id查询File
// @Tags File
// @Summary 用id查询File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.File true "用id查询File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /file/findFile [get]
func (fileApi *FileApi) FindFile(c *gin.Context) {
	var file file.File
	_ = c.ShouldBindQuery(&file)
	if refile, err := fileService.GetFile(file.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refile": refile}, c)
	}
}

// GetFileList 分页获取File列表
// @Tags File
// @Summary 分页获取File列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileSearch true "分页获取File列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /file/getFileList [get]
func (fileApi *FileApi) GetFileList(c *gin.Context) {
	var pageInfo fileReq.FileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := fileService.GetFileInfoList(pageInfo); err != nil {
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
