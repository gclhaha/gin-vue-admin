package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoHighlightRouter struct {}

// InitVideoHighlightRouter 初始化 视频高光 路由信息
func (s *VideoHighlightRouter) InitVideoHighlightRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	videoHighlightRouter := Router.Group("videoHighlight").Use(middleware.OperationRecord())
	videoHighlightRouterWithoutRecord := Router.Group("videoHighlight")
	videoHighlightRouterWithoutAuth := PublicRouter.Group("videoHighlight")
	{
		videoHighlightRouter.POST("createVideoHighlight", videoHighlightApi.CreateVideoHighlight)   // 新建视频高光
		videoHighlightRouter.DELETE("deleteVideoHighlight", videoHighlightApi.DeleteVideoHighlight) // 删除视频高光
		videoHighlightRouter.DELETE("deleteVideoHighlightByIds", videoHighlightApi.DeleteVideoHighlightByIds) // 批量删除视频高光
		videoHighlightRouter.PUT("updateVideoHighlight", videoHighlightApi.UpdateVideoHighlight)    // 更新视频高光
	}
	{
		videoHighlightRouterWithoutRecord.GET("findVideoHighlight", videoHighlightApi.FindVideoHighlight)        // 根据ID获取视频高光
		videoHighlightRouterWithoutRecord.GET("getVideoHighlightList", videoHighlightApi.GetVideoHighlightList)  // 获取视频高光列表
	}
	{
	    videoHighlightRouterWithoutAuth.GET("getVideoHighlightPublic", videoHighlightApi.GetVideoHighlightPublic)  // 视频高光开放接口
	}
}
