// 自动生成模板VenueItem
package leep

import (
	"strings"

	"github.com/google/uuid"
)

// 场馆场地 结构体  VenueItem
type VenueItem struct {
	Id           string `json:"id" form:"id" gorm:"primarykey;column:id;comment:场地ID，使用UUID;size:32;"`                    //场地ID，使用UUID
	VenueId      string `json:"venueId" form:"venueId" gorm:"column:venue_id;comment:所属场馆ID;size:32;" binding:"required"` //所属场馆ID
	ThumbnailUrl string `json:"thumbnailUrl" form:"thumbnailUrl" gorm:"column:thumbnail_url;comment:缩略图URL;"`             //缩略图URL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:场地名称;size:255;"`                               //场地名称
}

// TableName 场馆场地 VenueItem自定义表名 venue_item
func (VenueItem) TableName() string {
	return "venue_item"
}

// BeforeCreate 创建前添加id
func (venueItem *VenueItem) BeforeCreate() (err error) {
	venueItem.Id = strings.ReplaceAll(uuid.NewString(), "-", "")
	return
}
