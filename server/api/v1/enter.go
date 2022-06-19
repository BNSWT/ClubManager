package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/file"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/notification"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/task"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	ExampleApiGroup      example.ApiGroup
	NotificationApiGroup notification.ApiGroup
	TaskApiGroup         task.ApiGroup
	FileApiGroup         file.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
