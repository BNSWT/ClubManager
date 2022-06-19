import service from '@/utils/request'

// @Tags ExaFiles
// @Summary 创建ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaFiles true "创建ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaFiles/createExaFiles [post]
export const createExaFiles = (data) => {
  return service({
    url: '/exaFiles/createExaFiles',
    method: 'post',
    data
  })
}

// @Tags ExaFiles
// @Summary 删除ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaFiles true "删除ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaFiles/deleteExaFiles [delete]
export const deleteExaFiles = (data) => {
  return service({
    url: '/exaFiles/deleteExaFiles',
    method: 'delete',
    data
  })
}

// @Tags ExaFiles
// @Summary 删除ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaFiles/deleteExaFiles [delete]
export const deleteExaFilesByIds = (data) => {
  return service({
    url: '/exaFiles/deleteExaFilesByIds',
    method: 'delete',
    data
  })
}

// @Tags ExaFiles
// @Summary 更新ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaFiles true "更新ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exaFiles/updateExaFiles [put]
export const updateExaFiles = (data) => {
  return service({
    url: '/exaFiles/updateExaFiles',
    method: 'put',
    data
  })
}

// @Tags ExaFiles
// @Summary 用id查询ExaFiles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExaFiles true "用id查询ExaFiles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exaFiles/findExaFiles [get]
export const findExaFiles = (params) => {
  return service({
    url: '/exaFiles/findExaFiles',
    method: 'get',
    params
  })
}

// @Tags ExaFiles
// @Summary 分页获取ExaFiles列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExaFiles列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaFiles/getExaFilesList [get]
export const getExaFilesList = (params) => {
  return service({
    url: '/exaFiles/getExaFilesList',
    method: 'get',
    params
  })
}
