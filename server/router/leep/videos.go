package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideosRouter struct{}

func (s *VideosRouter) InitVideosRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	videosRouter := Router.Group("videos").Use(middleware.OperationRecord())
	videosRouterWithoutRecord := Router.Group("videos")
	videosRouterWithoutAuth := PublicRouter.Group("videos")
	{
		videosRouter.POST("createVideos", videosApi.CreateVideos)
		videosRouter.DELETE("deleteVideos", videosApi.DeleteVideos)
		videosRouter.DELETE("deleteVideosByIds", videosApi.DeleteVideosByIds)
		videosRouter.PUT("updateVideos", videosApi.UpdateVideos)
		videosRouter.POST("uploadVideos", videosApi.UploadVideos)
	}
	{
		videosRouterWithoutRecord.GET("findVideos", videosApi.FindVideos)
		videosRouterWithoutRecord.GET("getVideosList", videosApi.GetVideosList)
		videosRouterWithoutRecord.GET("loadVideo", videosApi.LoadVideo)
	}
	{
		videosRouterWithoutAuth.GET("getVideosPublic", videosApi.GetVideosPublic)
	}
}
