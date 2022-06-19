// 自动生成模板ExaFiles
package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExaFiles 结构体
// 如果含有time.Time 请自行import time包
type ExaFiles struct {
      global.GVA_MODEL
      FileName  string `json:"fileName" form:"fileName" gorm:"column:file_name;comment:;size:191;"`
      FileMd5  string `json:"fileMd5" form:"fileMd5" gorm:"column:file_md5;comment:;size:191;"`
      FilePath  string `json:"filePath" form:"filePath" gorm:"column:file_path;comment:;size:191;"`
      ChunkTotal  *int `json:"chunkTotal" form:"chunkTotal" gorm:"column:chunk_total;comment:;size:19;"`
      IsFinish  *bool `json:"isFinish" form:"isFinish" gorm:"column:is_finish;comment:;"`
}


// TableName ExaFiles 表名
func (ExaFiles) TableName() string {
  return "exa_files"
}

