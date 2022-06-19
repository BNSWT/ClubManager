package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/file"
	"github.com/flipped-aurora/gin-vue-admin/server/service/notfication"
	"github.com/flipped-aurora/gin-vue-admin/server/service/notification"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/task"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	NotficationServiceGroup  notfication.ServiceGroup
	NotificationServiceGroup notification.ServiceGroup
	TaskServiceGroup         task.ServiceGroup
	FileServiceGroup         file.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
