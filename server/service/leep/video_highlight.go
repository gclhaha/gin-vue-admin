package leep

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep"
    leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
)

type VideoHighlightService struct {}
// CreateVideoHighlight 创建视频高光记录
// Author [yourname](https://github.com/yourname)
func (videoHighlightService *VideoHighlightService) CreateVideoHighlight(videoHighlight *leep.VideoHighlight) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Create(videoHighlight).Error
	return err
}

// DeleteVideoHighlight 删除视频高光记录
// Author [yourname](https://github.com/yourname)
func (videoHighlightService *VideoHighlightService)DeleteVideoHighlight(id string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&leep.VideoHighlight{},"id = ?",id).Error
	return err
}

// DeleteVideoHighlightByIds 批量删除视频高光记录
// Author [yourname](https://github.com/yourname)
func (videoHighlightService *VideoHighlightService)DeleteVideoHighlightByIds(ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&[]leep.VideoHighlight{},"id in ?",ids).Error
	return err
}

// UpdateVideoHighlight 更新视频高光记录
// Author [yourname](https://github.com/yourname)
func (videoHighlightService *VideoHighlightService)UpdateVideoHighlight(videoHighlight leep.VideoHighlight) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Model(&leep.VideoHighlight{}).Where("id = ?",videoHighlight.Id).Updates(&videoHighlight).Error
	return err
}

// GetVideoHighlight 根据id获取视频高光记录
// Author [yourname](https://github.com/yourname)
func (videoHighlightService *VideoHighlightService)GetVideoHighlight(id string) (videoHighlight leep.VideoHighlight, err error) {
	err = global.MustGetGlobalDBByDBName("leep").Where("id = ?", id).First(&videoHighlight).Error
	return
}

// GetVideoHighlightInfoList 分页获取视频高光记录
// Author [yourname](https://github.com/yourname)
func (videoHighlightService *VideoHighlightService)GetVideoHighlightInfoList(info leepReq.VideoHighlightSearch) (list []leep.VideoHighlight, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("leep").Model(&leep.VideoHighlight{})
    var videoHighlights []leep.VideoHighlight
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&videoHighlights).Error
	return  videoHighlights, total, err
}
func (videoHighlightService *VideoHighlightService)GetVideoHighlightPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
