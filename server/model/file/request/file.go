package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FileSearch struct{
    file.File
    request.PageInfo
}
