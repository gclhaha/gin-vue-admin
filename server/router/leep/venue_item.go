package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VenueItemRouter struct {}

// InitVenueItemRouter 初始化 场馆场地 路由信息
func (s *VenueItemRouter) InitVenueItemRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	venueItemRouter := Router.Group("venueItem").Use(middleware.OperationRecord())
	venueItemRouterWithoutRecord := Router.Group("venueItem")
	venueItemRouterWithoutAuth := PublicRouter.Group("venueItem")
	{
		venueItemRouter.POST("createVenueItem", venueItemApi.CreateVenueItem)   // 新建场馆场地
		venueItemRouter.DELETE("deleteVenueItem", venueItemApi.DeleteVenueItem) // 删除场馆场地
		venueItemRouter.DELETE("deleteVenueItemByIds", venueItemApi.DeleteVenueItemByIds) // 批量删除场馆场地
		venueItemRouter.PUT("updateVenueItem", venueItemApi.UpdateVenueItem)    // 更新场馆场地
	}
	{
		venueItemRouterWithoutRecord.GET("findVenueItem", venueItemApi.FindVenueItem)        // 根据ID获取场馆场地
		venueItemRouterWithoutRecord.GET("getVenueItemList", venueItemApi.GetVenueItemList)  // 获取场馆场地列表
	}
	{
	    venueItemRouterWithoutAuth.GET("getVenueItemPublic", venueItemApi.GetVenueItemPublic)  // 场馆场地开放接口
	}
}
