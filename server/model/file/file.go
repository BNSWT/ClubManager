// 自动生成模板File
package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// File 结构体
// 如果含有time.Time 请自行import time包
type File struct {
      global.GVA_MODEL
      ParentID  *int `json:"parentID" form:"parentID" gorm:"column:parent_id;comment:;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      RealPath  string `json:"realPath" form:"realPath" gorm:"column:real_path;comment:;"`
}


// TableName File 表名
func (File) TableName() string {
  return "file"
}

