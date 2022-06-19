package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/notification"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type NotificationSearch struct{
    notification.Notification
    request.PageInfo
}
