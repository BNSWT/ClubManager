import service from '@/utils/request'

// @Tags Task
// @Summary 创建Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Task true "创建Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/createTask [post]
export const createTask = (data) => {
  return service({
    url: '/task/createTask',
    method: 'post',
    data
  })
}

// @Tags Task
// @Summary 删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Task true "删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTask [delete]
export const deleteTask = (data) => {
  return service({
    url: '/task/deleteTask',
    method: 'delete',
    data
  })
}

// @Tags Task
// @Summary 删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTask [delete]
export const deleteTaskByIds = (data) => {
  return service({
    url: '/task/deleteTaskByIds',
    method: 'delete',
    data
  })
}

// @Tags Task
// @Summary 更新Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Task true "更新Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateTask [put]
export const updateTask = (data) => {
  return service({
    url: '/task/updateTask',
    method: 'put',
    data
  })
}

// @Tags Task
// @Summary 用id查询Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Task true "用id查询Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTask [get]
export const findTask = (params) => {
  return service({
    url: '/task/findTask',
    method: 'get',
    params
  })
}

// @Tags Task
// @Summary 分页获取Task列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Task列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getTaskList [get]
export const getTaskList = (params) => {
  return service({
    url: '/task/getTaskList',
    method: 'get',
    params
  })
}
