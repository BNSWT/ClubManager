package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/file"
	"github.com/flipped-aurora/gin-vue-admin/server/router/notfication"
	"github.com/flipped-aurora/gin-vue-admin/server/router/notification"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/task"
)

type RouterGroup struct {
	System       system.RouterGroup
	Example      example.RouterGroup
	Notfication  notfication.RouterGroup
	Notification notification.RouterGroup
	Task         task.RouterGroup
	File         file.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
