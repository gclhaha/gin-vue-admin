package leep

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/leep"
    leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type VenueApi struct {}



// CreateVenue 创建场馆
// @Tags Venue
// @Summary 创建场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.Venue true "创建场馆"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /venue/createVenue [post]
func (venueApi *VenueApi) CreateVenue(c *gin.Context) {
	var venue leep.Venue
	err := c.ShouldBindJSON(&venue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = venueService.CreateVenue(&venue)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteVenue 删除场馆
// @Tags Venue
// @Summary 删除场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.Venue true "删除场馆"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /venue/deleteVenue [delete]
func (venueApi *VenueApi) DeleteVenue(c *gin.Context) {
	id := c.Query("id")
	err := venueService.DeleteVenue(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVenueByIds 批量删除场馆
// @Tags Venue
// @Summary 批量删除场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /venue/deleteVenueByIds [delete]
func (venueApi *VenueApi) DeleteVenueByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := venueService.DeleteVenueByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVenue 更新场馆
// @Tags Venue
// @Summary 更新场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body leep.Venue true "更新场馆"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /venue/updateVenue [put]
func (venueApi *VenueApi) UpdateVenue(c *gin.Context) {
	var venue leep.Venue
	err := c.ShouldBindJSON(&venue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = venueService.UpdateVenue(venue)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVenue 用id查询场馆
// @Tags Venue
// @Summary 用id查询场馆
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query leep.Venue true "用id查询场馆"
// @Success 200 {object} response.Response{data=leep.Venue,msg=string} "查询成功"
// @Router /venue/findVenue [get]
func (venueApi *VenueApi) FindVenue(c *gin.Context) {
	id := c.Query("id")
	revenue, err := venueService.GetVenue(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(revenue, c)
}

// GetVenueList 分页获取场馆列表
// @Tags Venue
// @Summary 分页获取场馆列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VenueSearch true "分页获取场馆列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /venue/getVenueList [get]
func (venueApi *VenueApi) GetVenueList(c *gin.Context) {
	var pageInfo leepReq.VenueSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := venueService.GetVenueInfoList(pageInfo)
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

// GetVenuePublic 不需要鉴权的场馆接口
// @Tags Venue
// @Summary 不需要鉴权的场馆接口
// @accept application/json
// @Produce application/json
// @Param data query leepReq.VenueSearch true "分页获取场馆列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /venue/getVenuePublic [get]
func (venueApi *VenueApi) GetVenuePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    venueService.GetVenuePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的场馆接口信息",
    }, "获取成功", c)
}
