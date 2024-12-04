package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep/response"
	leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
)

type VenueItemService struct {}
// CreateVenueItem 创建场馆场地记录
// Author [yourname](https://github.com/yourname)
func (venueItemService *VenueItemService) CreateVenueItem(venueItem *leep.VenueItem) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Create(venueItem).Error
	return err
}

// DeleteVenueItem 删除场馆场地记录
// Author [yourname](https://github.com/yourname)
func (venueItemService *VenueItemService)DeleteVenueItem(id string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&leep.VenueItem{},"id = ?",id).Error
	return err
}

// DeleteVenueItemByIds 批量删除场馆场地记录
// Author [yourname](https://github.com/yourname)
func (venueItemService *VenueItemService)DeleteVenueItemByIds(ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&[]leep.VenueItem{},"id in ?",ids).Error
	return err
}

// UpdateVenueItem 更新场馆场地记录
// Author [yourname](https://github.com/yourname)
func (venueItemService *VenueItemService)UpdateVenueItem(venueItem leep.VenueItem) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Model(&leep.VenueItem{}).Where("id = ?",venueItem.Id).Updates(&venueItem).Error
	return err
}

// GetVenueItem 根据id获取场馆场地记录
// Author [yourname](https://github.com/yourname)
func (venueItemService *VenueItemService)GetVenueItem(id string) (venueItem leep.VenueItem, err error) {
	err = global.MustGetGlobalDBByDBName("leep").Where("id = ?", id).First(&venueItem).Error
	return
}

// GetVenueItemInfoList 分页获取场馆场地记录
// Author [yourname](https://github.com/yourname)
func (venueItemService *VenueItemService)GetVenueItemInfoList(info leepReq.VenueItemSearch) (list []response.VenueItemDto, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("leep").Model(&leep.VenueItem{})
	var venueItems []leep.VenueItem
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&venueItems).Error
	if err != nil {
		return
	}

	// 根据venueItems的venue_id获取venue的信息
	var venueIds []string
	for _, item := range venueItems {
		venueIds = append(venueIds, item.VenueId)
	}

	venueService := VenueService{}
	venues, err := venueService.GetVenueInfoByIds(venueIds)
	if err != nil {
		return
	}

	venueMap := make(map[string]string)
	for _, venue := range venues {
		venueMap[venue.Id] = venue.Name
	}

	for _, item := range venueItems {
		dto := response.ToVenueItemDto(item)
		dto.VenueName = venueMap[item.VenueId]
		list = append(list, dto)
	}

	return list, total, err
}

func (venueItemService *VenueItemService)GetVenueItemPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
