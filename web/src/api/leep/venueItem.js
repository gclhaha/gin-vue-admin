import service from '@/utils/request'
// @Tags VenueItem
// @Summary 创建场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VenueItem true "创建场馆场地"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /venueItem/createVenueItem [post]
export const createVenueItem = (data) => {
  return service({
    url: '/venueItem/createVenueItem',
    method: 'post',
    data
  })
}

// @Tags VenueItem
// @Summary 删除场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VenueItem true "删除场馆场地"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /venueItem/deleteVenueItem [delete]
export const deleteVenueItem = (params) => {
  return service({
    url: '/venueItem/deleteVenueItem',
    method: 'delete',
    params
  })
}

// @Tags VenueItem
// @Summary 批量删除场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除场馆场地"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /venueItem/deleteVenueItem [delete]
export const deleteVenueItemByIds = (params) => {
  return service({
    url: '/venueItem/deleteVenueItemByIds',
    method: 'delete',
    params
  })
}

// @Tags VenueItem
// @Summary 更新场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VenueItem true "更新场馆场地"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /venueItem/updateVenueItem [put]
export const updateVenueItem = (data) => {
  return service({
    url: '/venueItem/updateVenueItem',
    method: 'put',
    data
  })
}

// @Tags VenueItem
// @Summary 用id查询场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VenueItem true "用id查询场馆场地"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /venueItem/findVenueItem [get]
export const findVenueItem = (params) => {
  return service({
    url: '/venueItem/findVenueItem',
    method: 'get',
    params
  })
}

// @Tags VenueItem
// @Summary 分页获取场馆场地列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取场馆场地列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /venueItem/getVenueItemList [get]
export const getVenueItemList = (params) => {
  return service({
    url: '/venueItem/getVenueItemList',
    method: 'get',
    params
  })
}

// @Tags VenueItem
// @Summary 不需要鉴权的场馆场地接口
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VenueItemSearch true "分页获取场馆场地列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /venueItem/getVenueItemPublic [get]
export const getVenueItemPublic = () => {
  return service({
    url: '/venueItem/getVenueItemPublic',
    method: 'get',
  })
}
