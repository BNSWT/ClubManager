package notification

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/notification"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    notificationReq "github.com/flipped-aurora/gin-vue-admin/server/model/notification/request"
)

type NotificationService struct {
}

// CreateNotification 创建Notification记录
// Author [piexlmax](https://github.com/piexlmax)
func (notiService *NotificationService) CreateNotification(noti notification.Notification) (err error) {
	err = global.GVA_DB.Create(&noti).Error
	return err
}

// DeleteNotification 删除Notification记录
// Author [piexlmax](https://github.com/piexlmax)
func (notiService *NotificationService)DeleteNotification(noti notification.Notification) (err error) {
	err = global.GVA_DB.Delete(&noti).Error
	return err
}

// DeleteNotificationByIds 批量删除Notification记录
// Author [piexlmax](https://github.com/piexlmax)
func (notiService *NotificationService)DeleteNotificationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]notification.Notification{},"id in ?",ids.Ids).Error
	return err
}

// UpdateNotification 更新Notification记录
// Author [piexlmax](https://github.com/piexlmax)
func (notiService *NotificationService)UpdateNotification(noti notification.Notification) (err error) {
	err = global.GVA_DB.Save(&noti).Error
	return err
}

// GetNotification 根据id获取Notification记录
// Author [piexlmax](https://github.com/piexlmax)
func (notiService *NotificationService)GetNotification(id uint) (noti notification.Notification, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&noti).Error
	return
}

// GetNotificationInfoList 分页获取Notification记录
// Author [piexlmax](https://github.com/piexlmax)
func (notiService *NotificationService)GetNotificationInfoList(info notificationReq.NotificationSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&notification.Notification{})
    var notis []notification.Notification
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&notis).Error
	return  notis, total, err
}
