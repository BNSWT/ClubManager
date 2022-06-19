package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type ExaFilesService struct {
}

// CreateExaFiles 创建ExaFiles记录
// Author [piexlmax](https://github.com/piexlmax)
func (exaFilesService *ExaFilesService) CreateExaFiles(exaFiles file.ExaFiles) (err error) {
	err = global.GVA_DB.Create(&exaFiles).Error
	return err
}

// DeleteExaFiles 删除ExaFiles记录
// Author [piexlmax](https://github.com/piexlmax)
func (exaFilesService *ExaFilesService)DeleteExaFiles(exaFiles file.ExaFiles) (err error) {
	err = global.GVA_DB.Delete(&exaFiles).Error
	return err
}

// DeleteExaFilesByIds 批量删除ExaFiles记录
// Author [piexlmax](https://github.com/piexlmax)
func (exaFilesService *ExaFilesService)DeleteExaFilesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.ExaFiles{},"id in ?",ids.Ids).Error
	return err
}

// UpdateExaFiles 更新ExaFiles记录
// Author [piexlmax](https://github.com/piexlmax)
func (exaFilesService *ExaFilesService)UpdateExaFiles(exaFiles file.ExaFiles) (err error) {
	err = global.GVA_DB.Save(&exaFiles).Error
	return err
}

// GetExaFiles 根据id获取ExaFiles记录
// Author [piexlmax](https://github.com/piexlmax)
func (exaFilesService *ExaFilesService)GetExaFiles(id uint) (exaFiles file.ExaFiles, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&exaFiles).Error
	return
}

// GetExaFilesInfoList 分页获取ExaFiles记录
// Author [piexlmax](https://github.com/piexlmax)
func (exaFilesService *ExaFilesService)GetExaFilesInfoList(info fileReq.ExaFilesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&file.ExaFiles{})
    var exaFiless []file.ExaFiles
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&exaFiless).Error
	return  exaFiless, total, err
}
