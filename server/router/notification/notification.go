package notification

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NotificationRouter struct {
}

// InitNotificationRouter 初始化 Notification 路由信息
func (s *NotificationRouter) InitNotificationRouter(Router *gin.RouterGroup) {
	notiRouter := Router.Group("noti").Use(middleware.OperationRecord())
	notiRouterWithoutRecord := Router.Group("noti")
	var notiApi = v1.ApiGroupApp.NotificationApiGroup.NotificationApi
	{
		notiRouter.POST("createNotification", notiApi.CreateNotification)   // 新建Notification
		notiRouter.DELETE("deleteNotification", notiApi.DeleteNotification) // 删除Notification
		notiRouter.DELETE("deleteNotificationByIds", notiApi.DeleteNotificationByIds) // 批量删除Notification
		notiRouter.PUT("updateNotification", notiApi.UpdateNotification)    // 更新Notification
	}
	{
		notiRouterWithoutRecord.GET("findNotification", notiApi.FindNotification)        // 根据ID获取Notification
		notiRouterWithoutRecord.GET("getNotificationList", notiApi.GetNotificationList)  // 获取Notification列表
	}
}
