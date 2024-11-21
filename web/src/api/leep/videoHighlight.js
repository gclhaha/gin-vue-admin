import service from '@/utils/request'
// @Tags VideoHighlight
// @Summary 创建视频高光
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoHighlight true "创建视频高光"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /videoHighlight/createVideoHighlight [post]
export const createVideoHighlight = (data) => {
  return service({
    url: '/videoHighlight/createVideoHighlight',
    method: 'post',
    data
  })
}

// @Tags VideoHighlight
// @Summary 删除视频高光
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoHighlight true "删除视频高光"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoHighlight/deleteVideoHighlight [delete]
export const deleteVideoHighlight = (params) => {
  return service({
    url: '/videoHighlight/deleteVideoHighlight',
    method: 'delete',
    params
  })
}

// @Tags VideoHighlight
// @Summary 批量删除视频高光
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除视频高光"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoHighlight/deleteVideoHighlight [delete]
export const deleteVideoHighlightByIds = (params) => {
  return service({
    url: '/videoHighlight/deleteVideoHighlightByIds',
    method: 'delete',
    params
  })
}

// @Tags VideoHighlight
// @Summary 更新视频高光
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoHighlight true "更新视频高光"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoHighlight/updateVideoHighlight [put]
export const updateVideoHighlight = (data) => {
  return service({
    url: '/videoHighlight/updateVideoHighlight',
    method: 'put',
    data
  })
}

// @Tags VideoHighlight
// @Summary 用id查询视频高光
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VideoHighlight true "用id查询视频高光"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoHighlight/findVideoHighlight [get]
export const findVideoHighlight = (params) => {
  return service({
    url: '/videoHighlight/findVideoHighlight',
    method: 'get',
    params
  })
}

// @Tags VideoHighlight
// @Summary 分页获取视频高光列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取视频高光列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoHighlight/getVideoHighlightList [get]
export const getVideoHighlightList = (params) => {
  return service({
    url: '/videoHighlight/getVideoHighlightList',
    method: 'get',
    params
  })
}

// @Tags VideoHighlight
// @Summary 不需要鉴权的视频高光接口
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VideoHighlightSearch true "分页获取视频高光列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videoHighlight/getVideoHighlightPublic [get]
export const getVideoHighlightPublic = () => {
  return service({
    url: '/videoHighlight/getVideoHighlightPublic',
    method: 'get',
  })
}
