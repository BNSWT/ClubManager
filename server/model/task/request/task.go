package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TaskSearch struct{
    task.Task
    request.PageInfo
}
