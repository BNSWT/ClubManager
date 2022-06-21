// 自动生成模板Notification
package notification

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Notification 结构体
// 如果含有time.Time 请自行import time包
type Notification struct {
	global.GVA_MODEL
	Title          string `json:"title" form:"title" gorm:"column:title;comment:;"`
	Content        string `json:"content" form:"content" gorm:"column:content;comment:;"`
	AttachmentPath string `json:"attachmentPath" form:"attachment" gorm:"column:attachment;comment:附件路径;"`
	StickyOnTop    *bool  `json:"stickyOnTop" form:"stickyOnTop" gorm:"column:sticky_on_top;comment:通知是否置顶;"`
	Coverage       string `json:"coverage" form:"coverage" gorm:"column:coverage;type:enum('所有成员', '活动部', '学术部', '宣传部');comment:;"`
}

// TableName Notification 表名
func (Notification) TableName() string {
	return "notification"
}
