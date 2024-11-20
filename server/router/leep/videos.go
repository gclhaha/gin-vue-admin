package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideosRouter struct {}

// InitVideosRouter 初始化 视频管理 路由信息
func (s *VideosRouter) InitVideosRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	videosRouter := Router.Group("videos").Use(middleware.OperationRecord())
	videosRouterWithoutRecord := Router.Group("videos")
	videosRouterWithoutAuth := PublicRouter.Group("videos")
	{
		videosRouter.POST("createVideos", videosApi.CreateVideos)   // 新建视频管理
		videosRouter.DELETE("deleteVideos", videosApi.DeleteVideos) // 删除视频管理
		videosRouter.DELETE("deleteVideosByIds", videosApi.DeleteVideosByIds) // 批量删除视频管理
		videosRouter.PUT("updateVideos", videosApi.UpdateVideos)    // 更新视频管理
		videosRouter.POST("uploadVideos", videosApi.UploadVideos)	// 上传视频管理
	}
	{
		videosRouterWithoutRecord.GET("findVideos", videosApi.FindVideos)        // 根据ID获取视频管理
		videosRouterWithoutRecord.GET("getVideosList", videosApi.GetVideosList)  // 获取视频管理列表
	}
	{
	    videosRouterWithoutAuth.GET("getVideosPublic", videosApi.GetVideosPublic)  // 视频管理开放接口
	}
}
