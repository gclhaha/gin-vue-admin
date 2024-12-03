import service from '@/utils/request'
// @Tags Venue
// @Summary 创建场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Venue true "创建场馆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /venue/createVenue [post]
export const createVenue = (data) => {
  return service({
    url: '/venue/createVenue',
    method: 'post',
    data
  })
}

// @Tags Venue
// @Summary 删除场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Venue true "删除场馆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /venue/deleteVenue [delete]
export const deleteVenue = (params) => {
  return service({
    url: '/venue/deleteVenue',
    method: 'delete',
    params
  })
}

// @Tags Venue
// @Summary 批量删除场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除场馆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /venue/deleteVenue [delete]
export const deleteVenueByIds = (params) => {
  return service({
    url: '/venue/deleteVenueByIds',
    method: 'delete',
    params
  })
}

// @Tags Venue
// @Summary 更新场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Venue true "更新场馆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /venue/updateVenue [put]
export const updateVenue = (data) => {
  return service({
    url: '/venue/updateVenue',
    method: 'put',
    data
  })
}

// @Tags Venue
// @Summary 用id查询场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Venue true "用id查询场馆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /venue/findVenue [get]
export const findVenue = (params) => {
  return service({
    url: '/venue/findVenue',
    method: 'get',
    params
  })
}

// @Tags Venue
// @Summary 分页获取场馆列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取场馆列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /venue/getVenueList [get]
export const getVenueList = (params) => {
  return service({
    url: '/venue/getVenueList',
    method: 'get',
    params
  })
}

// @Tags Venue
// @Summary 不需要鉴权的场馆接口
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VenueSearch true "分页获取场馆列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /venue/getVenuePublic [get]
export const getVenuePublic = () => {
  return service({
    url: '/venue/getVenuePublic',
    method: 'get',
  })
}
