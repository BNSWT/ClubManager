package notification

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/notification"
	notificationReq "github.com/flipped-aurora/gin-vue-admin/server/model/notification/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NotificationApi struct {
}

var notiService = service.ServiceGroupApp.NotificationServiceGroup.NotificationService


// CreateNotification 创建Notification
// @Tags Notification
// @Summary 创建Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body notification.Notification true "创建Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /noti/createNotification [post]
func (notiApi *NotificationApi) CreateNotification(c *gin.Context) {
	var noti notification.Notification
	_ = c.ShouldBindJSON(&noti)
	if err := notiService.CreateNotification(noti); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNotification 删除Notification
// @Tags Notification
// @Summary 删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body notification.Notification true "删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /noti/deleteNotification [delete]
func (notiApi *NotificationApi) DeleteNotification(c *gin.Context) {
	var noti notification.Notification
	_ = c.ShouldBindJSON(&noti)
	if err := notiService.DeleteNotification(noti); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNotificationByIds 批量删除Notification
// @Tags Notification
// @Summary 批量删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /noti/deleteNotificationByIds [delete]
func (notiApi *NotificationApi) DeleteNotificationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := notiService.DeleteNotificationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNotification 更新Notification
// @Tags Notification
// @Summary 更新Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body notification.Notification true "更新Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /noti/updateNotification [put]
func (notiApi *NotificationApi) UpdateNotification(c *gin.Context) {
	var noti notification.Notification
	_ = c.ShouldBindJSON(&noti)
	if err := notiService.UpdateNotification(noti); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNotification 用id查询Notification
// @Tags Notification
// @Summary 用id查询Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query notification.Notification true "用id查询Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /noti/findNotification [get]
func (notiApi *NotificationApi) FindNotification(c *gin.Context) {
	var noti notification.Notification
	_ = c.ShouldBindQuery(&noti)
	if renoti, err := notiService.GetNotification(noti.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renoti": renoti}, c)
	}
}

// GetNotificationList 分页获取Notification列表
// @Tags Notification
// @Summary 分页获取Notification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query notificationReq.NotificationSearch true "分页获取Notification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /noti/getNotificationList [get]
func (notiApi *NotificationApi) GetNotificationList(c *gin.Context) {
	var pageInfo notificationReq.NotificationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := notiService.GetNotificationInfoList(pageInfo); err != nil {
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
