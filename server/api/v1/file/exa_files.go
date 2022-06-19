package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/file"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ExaFilesApi struct {
}

var exaFilesService = service.ServiceGroupApp.FileServiceGroup.ExaFilesService


// CreateExaFiles 创建ExaFiles
// @Tags ExaFiles
// @Summary 创建ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.ExaFiles true "创建ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaFiles/createExaFiles [post]
func (exaFilesApi *ExaFilesApi) CreateExaFiles(c *gin.Context) {
	var exaFiles file.ExaFiles
	_ = c.ShouldBindJSON(&exaFiles)
	if err := exaFilesService.CreateExaFiles(exaFiles); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExaFiles 删除ExaFiles
// @Tags ExaFiles
// @Summary 删除ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.ExaFiles true "删除ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaFiles/deleteExaFiles [delete]
func (exaFilesApi *ExaFilesApi) DeleteExaFiles(c *gin.Context) {
	var exaFiles file.ExaFiles
	_ = c.ShouldBindJSON(&exaFiles)
	if err := exaFilesService.DeleteExaFiles(exaFiles); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExaFilesByIds 批量删除ExaFiles
// @Tags ExaFiles
// @Summary 批量删除ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exaFiles/deleteExaFilesByIds [delete]
func (exaFilesApi *ExaFilesApi) DeleteExaFilesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := exaFilesService.DeleteExaFilesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExaFiles 更新ExaFiles
// @Tags ExaFiles
// @Summary 更新ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.ExaFiles true "更新ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exaFiles/updateExaFiles [put]
func (exaFilesApi *ExaFilesApi) UpdateExaFiles(c *gin.Context) {
	var exaFiles file.ExaFiles
	_ = c.ShouldBindJSON(&exaFiles)
	if err := exaFilesService.UpdateExaFiles(exaFiles); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExaFiles 用id查询ExaFiles
// @Tags ExaFiles
// @Summary 用id查询ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.ExaFiles true "用id查询ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exaFiles/findExaFiles [get]
func (exaFilesApi *ExaFilesApi) FindExaFiles(c *gin.Context) {
	var exaFiles file.ExaFiles
	_ = c.ShouldBindQuery(&exaFiles)
	if reexaFiles, err := exaFilesService.GetExaFiles(exaFiles.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexaFiles": reexaFiles}, c)
	}
}

// GetExaFilesList 分页获取ExaFiles列表
// @Tags ExaFiles
// @Summary 分页获取ExaFiles列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.ExaFilesSearch true "分页获取ExaFiles列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaFiles/getExaFilesList [get]
func (exaFilesApi *ExaFilesApi) GetExaFilesList(c *gin.Context) {
	var pageInfo fileReq.ExaFilesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := exaFilesService.GetExaFilesInfoList(pageInfo); err != nil {
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
