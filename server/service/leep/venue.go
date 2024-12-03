package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep"
    leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
)

type VenueService struct {}
// CreateVenue 创建场馆记录
// Author [yourname](https://github.com/yourname)
func (venueService *VenueService) CreateVenue(venue *leep.Venue) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Create(venue).Error
	return err
}

// DeleteVenue 删除场馆记录
// Author [yourname](https://github.com/yourname)
func (venueService *VenueService)DeleteVenue(id string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&leep.Venue{},"id = ?",id).Error
	return err
}

// DeleteVenueByIds 批量删除场馆记录
// Author [yourname](https://github.com/yourname)
func (venueService *VenueService)DeleteVenueByIds(ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&[]leep.Venue{},"id in ?",ids).Error
	return err
}

// UpdateVenue 更新场馆记录
// Author [yourname](https://github.com/yourname)
func (venueService *VenueService)UpdateVenue(venue leep.Venue) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Model(&leep.Venue{}).Where("id = ?",venue.Id).Updates(&venue).Error
	return err
}

// GetVenue 根据id获取场馆记录
// Author [yourname](https://github.com/yourname)
func (venueService *VenueService)GetVenue(id string) (venue leep.Venue, err error) {
	err = global.MustGetGlobalDBByDBName("leep").Where("id = ?", id).First(&venue).Error
	return
}

// GetVenueInfoList 分页获取场馆记录
// Author [yourname](https://github.com/yourname)
func (venueService *VenueService)GetVenueInfoList(info leepReq.VenueSearch) (list []leep.Venue, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("leep").Model(&leep.Venue{})
    var venues []leep.Venue
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&venues).Error
	return  venues, total, err
}
func (venueService *VenueService)GetVenuePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
