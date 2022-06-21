package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type FileService struct {
}

// CreateFile 创建File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) CreateFile(file file.File) (err error) {
	err = global.GVA_DB.Create(&file).Error
	return err
}

// DeleteFile 删除File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) DeleteFile(file file.File) (err error) {
	err = global.GVA_DB.Delete(&file).Error
	return err
}

// DeleteFileByIds 批量删除File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) DeleteFileByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.File{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFile 更新File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) UpdateFile(file file.File) (err error) {
	err = global.GVA_DB.Save(&file).Error
	return err
}

// GetFile 根据id获取File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) GetFile(id uint) (file file.File, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&file).Error
	return
}

// GetFileInfoList 分页获取File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) GetFileInfoList(info fileReq.FileSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&file.File{})
	var files []file.File
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&files).Error
	return files, total, err
}
