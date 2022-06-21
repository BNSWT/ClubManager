import service from '@/utils/request'

// @Tags File
// @Summary 创建File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.File true "创建File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /file/createFile [post]
export const createFile = (data) => {
  return service({
    url: '/file/createFile',
    method: 'post',
    data
  })
}

// @Tags File
// @Summary 删除File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.File true "删除File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /file/deleteFile [delete]
export const deleteFile = (data) => {
  return service({
    url: '/file/deleteFile',
    method: 'delete',
    data
  })
}

// @Tags File
// @Summary 删除File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /file/deleteFile [delete]
export const deleteFileByIds = (data) => {
  return service({
    url: '/file/deleteFileByIds',
    method: 'delete',
    data
  })
}

// @Tags File
// @Summary 更新File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.File true "更新File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /file/updateFile [put]
export const updateFile = (data) => {
  return service({
    url: '/file/updateFile',
    method: 'put',
    data
  })
}

// @Tags File
// @Summary 用id查询File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.File true "用id查询File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /file/findFile [get]
export const findFile = (params) => {
  return service({
    url: '/file/findFile',
    method: 'get',
    params
  })
}

// @Tags File
// @Summary 分页获取File列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取File列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /file/getFileList [get]
export const getFileList = (params) => {
  return service({
    url: '/file/getFileList',
    method: 'get',
    params
  })
}
