import service from '@/utils/request'
// @Tags Videos
// @Summary 创建视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Videos true "创建视频管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /videos/createVideos [post]
export const createVideos = (data) => {
  return service({
    url: '/videos/createVideos',
    method: 'post',
    data
  })
}

// @Tags Videos
// @Summary 删除视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Videos true "删除视频管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videos/deleteVideos [delete]
export const deleteVideos = (params) => {
  return service({
    url: '/videos/deleteVideos',
    method: 'delete',
    params
  })
}

// @Tags Videos
// @Summary 批量删除视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除视频管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videos/deleteVideos [delete]
export const deleteVideosByIds = (params) => {
  return service({
    url: '/videos/deleteVideosByIds',
    method: 'delete',
    params
  })
}

// @Tags Videos
// @Summary 更新视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Videos true "更新视频管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videos/updateVideos [put]
export const updateVideos = (data) => {
  return service({
    url: '/videos/updateVideos',
    method: 'put',
    data
  })
}

// @Tags Videos
// @Summary 用id查询视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Videos true "用id查询视频管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videos/findVideos [get]
export const findVideos = (params) => {
  return service({
    url: '/videos/findVideos',
    method: 'get',
    params
  })
}

// @Tags Videos
// @Summary 分页获取视频管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取视频管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videos/getVideosList [get]
export const getVideosList = (params) => {
  return service({
    url: '/videos/getVideosList',
    method: 'get',
    params
  })
}

// @Tags Videos
// @Summary 不需要鉴权的视频管理接口
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VideosSearch true "分页获取视频管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videos/getVideosPublic [get]
export const getVideosPublic = () => {
  return service({
    url: '/videos/getVideosPublic',
    method: 'get',
  })
}

// @Tags Videos
// @Summary 上传视频
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce application/json
// @Param userId formData string true "用户ID"
// @Param file formData file true "视频文件"
// @Success 200 {string} string "{"success":true,"data":{videoUrl: string},"msg":"上传成功"}"
// @Router /videos/uploadVideos [post]
export const uploadVideos = (data) => {
  return service({
    url: '/videos/uploadVideos',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'multipart/form-data' // 重要：设置 Content-Type 为 multipart/form-data
    }
  })
}