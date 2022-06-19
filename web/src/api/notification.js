import service from '@/utils/request'

// @Tags Notification
// @Summary 创建Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "创建Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /noti/createNotification [post]
export const createNotification = (data) => {
  return service({
    url: '/noti/createNotification',
    method: 'post',
    data
  })
}

// @Tags Notification
// @Summary 删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /noti/deleteNotification [delete]
export const deleteNotification = (data) => {
  return service({
    url: '/noti/deleteNotification',
    method: 'delete',
    data
  })
}

// @Tags Notification
// @Summary 删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /noti/deleteNotification [delete]
export const deleteNotificationByIds = (data) => {
  return service({
    url: '/noti/deleteNotificationByIds',
    method: 'delete',
    data
  })
}

// @Tags Notification
// @Summary 更新Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "更新Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /noti/updateNotification [put]
export const updateNotification = (data) => {
  return service({
    url: '/noti/updateNotification',
    method: 'put',
    data
  })
}

// @Tags Notification
// @Summary 用id查询Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Notification true "用id查询Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /noti/findNotification [get]
export const findNotification = (params) => {
  return service({
    url: '/noti/findNotification',
    method: 'get',
    params
  })
}

// @Tags Notification
// @Summary 分页获取Notification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Notification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /noti/getNotificationList [get]
export const getNotificationList = (params) => {
  return service({
    url: '/noti/getNotificationList',
    method: 'get',
    params
  })
}
