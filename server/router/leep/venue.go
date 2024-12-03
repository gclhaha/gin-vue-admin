package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VenueRouter struct {}

// InitVenueRouter 初始化 场馆 路由信息
func (s *VenueRouter) InitVenueRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	venueRouter := Router.Group("venue").Use(middleware.OperationRecord())
	venueRouterWithoutRecord := Router.Group("venue")
	venueRouterWithoutAuth := PublicRouter.Group("venue")
	{
		venueRouter.POST("createVenue", venueApi.CreateVenue)   // 新建场馆
		venueRouter.DELETE("deleteVenue", venueApi.DeleteVenue) // 删除场馆
		venueRouter.DELETE("deleteVenueByIds", venueApi.DeleteVenueByIds) // 批量删除场馆
		venueRouter.PUT("updateVenue", venueApi.UpdateVenue)    // 更新场馆
	}
	{
		venueRouterWithoutRecord.GET("findVenue", venueApi.FindVenue)        // 根据ID获取场馆
		venueRouterWithoutRecord.GET("getVenueList", venueApi.GetVenueList)  // 获取场馆列表
	}
	{
	    venueRouterWithoutAuth.GET("getVenuePublic", venueApi.GetVenuePublic)  // 场馆开放接口
	}
}
