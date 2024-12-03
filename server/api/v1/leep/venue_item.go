package leep

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/leep"
    leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type VenueItemApi struct {}



// CreateVenueItem 创建场馆场地
// @Tags VenueItem
// @Summary 创建场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.VenueItem true "创建场馆场地"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /venueItem/createVenueItem [post]
func (venueItemApi *VenueItemApi) CreateVenueItem(c *gin.Context) {
	var venueItem leep.VenueItem
	err := c.ShouldBindJSON(&venueItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = venueItemService.CreateVenueItem(&venueItem)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteVenueItem 删除场馆场地
// @Tags VenueItem
// @Summary 删除场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.VenueItem true "删除场馆场地"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /venueItem/deleteVenueItem [delete]
func (venueItemApi *VenueItemApi) DeleteVenueItem(c *gin.Context) {
	id := c.Query("id")
	err := venueItemService.DeleteVenueItem(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVenueItemByIds 批量删除场馆场地
// @Tags VenueItem
// @Summary 批量删除场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /venueItem/deleteVenueItemByIds [delete]
func (venueItemApi *VenueItemApi) DeleteVenueItemByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := venueItemService.DeleteVenueItemByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVenueItem 更新场馆场地
// @Tags VenueItem
// @Summary 更新场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.VenueItem true "更新场馆场地"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /venueItem/updateVenueItem [put]
func (venueItemApi *VenueItemApi) UpdateVenueItem(c *gin.Context) {
	var venueItem leep.VenueItem
	err := c.ShouldBindJSON(&venueItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = venueItemService.UpdateVenueItem(venueItem)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVenueItem 用id查询场馆场地
// @Tags VenueItem
// @Summary 用id查询场馆场地
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query leep.VenueItem true "用id查询场馆场地"
// @Success 200 {object} response.Response{data=leep.VenueItem,msg=string} "查询成功"
// @Router /venueItem/findVenueItem [get]
func (venueItemApi *VenueItemApi) FindVenueItem(c *gin.Context) {
	id := c.Query("id")
	revenueItem, err := venueItemService.GetVenueItem(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(revenueItem, c)
}

// GetVenueItemList 分页获取场馆场地列表
// @Tags VenueItem
// @Summary 分页获取场馆场地列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VenueItemSearch true "分页获取场馆场地列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /venueItem/getVenueItemList [get]
func (venueItemApi *VenueItemApi) GetVenueItemList(c *gin.Context) {
	var pageInfo leepReq.VenueItemSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := venueItemService.GetVenueItemInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetVenueItemPublic 不需要鉴权的场馆场地接口
// @Tags VenueItem
// @Summary 不需要鉴权的场馆场地接口
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VenueItemSearch true "分页获取场馆场地列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /venueItem/getVenueItemPublic [get]
func (venueItemApi *VenueItemApi) GetVenueItemPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    venueItemService.GetVenueItemPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的场馆场地接口信息",
    }, "获取成功", c)
}
