package leep

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep"
	leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideosApi struct{}

// CreateVideos 创建视频管理
// @Tags Videos
// @Summary 创建视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.Videos true "创建视频管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /videos/createVideos [post]
func (videosApi *VideosApi) CreateVideos(c *gin.Context) {
	userId := c.PostForm("userId")
	videoId := c.PostForm("videoId")
	file, err := c.FormFile("file")

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("文件上传失败: %s", err.Error()), c)
		return
	}

	if userId == "" { // Ensure userId is provided
		response.FailWithMessage("用户ID不能为空", c)
		return
	}

	// Call the service function to handle the upload and database interaction
	videoUrl, err := videosService.CreateVideos(userId, videoId, file)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithData(gin.H{"videoUrl": videoUrl}, c) // Return the videoUrl
}

// DeleteVideos 删除视频管理
// @Tags Videos
// @Summary 删除视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.Videos true "删除视频管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /videos/deleteVideos [delete]
func (videosApi *VideosApi) DeleteVideos(c *gin.Context) {
	id := c.Query("id")
	err := videosService.DeleteVideos(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVideosByIds 批量删除视频管理
// @Tags Videos
// @Summary 批量删除视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /videos/deleteVideosByIds [delete]
func (videosApi *VideosApi) DeleteVideosByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := videosService.DeleteVideosByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVideos 更新视频管理
// @Tags Videos
// @Summary 更新视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.Videos true "更新视频管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /videos/updateVideos [put]
func (videosApi *VideosApi) UpdateVideos(c *gin.Context) {
	var videos leep.Videos
	err := c.ShouldBindJSON(&videos)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = videosService.UpdateVideos(videos)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVideos 用id查询视频管理
// @Tags Videos
// @Summary 用id查询视频管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query leep.Videos true "用id查询视频管理"
// @Success 200 {object} response.Response{data=leep.Videos,msg=string} "查询成功"
// @Router /videos/findVideos [get]
func (videosApi *VideosApi) FindVideos(c *gin.Context) {
	id := c.Query("id")
	revideos, err := videosService.GetVideos(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(revideos, c)
}

// GetVideosList 分页获取视频管理列表
// @Tags Videos
// @Summary 分页获取视频管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VideosSearch true "分页获取视频管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /videos/getVideosList [get]
func (videosApi *VideosApi) GetVideosList(c *gin.Context) {
	var pageInfo leepReq.VideosSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := videosService.GetVideosInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetVideosPublic 不需要鉴权的视频管理接口
// @Tags Videos
// @Summary 不需要鉴权的视频管理接口
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VideosSearch true "分页获取视频管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videos/getVideosPublic [get]
func (videosApi *VideosApi) GetVideosPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	videosService.GetVideosPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的视频管理接口信息",
	}, "获取成功", c)
}
