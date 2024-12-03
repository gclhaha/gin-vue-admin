// 自动生成模板Venue
package leep

import (
	"strings"

	"github.com/google/uuid"
)

// 场馆 结构体  Venue
type Venue struct {
    Id  string `json:"id" form:"id" gorm:"primarykey;column:id;comment:场馆ID，使用UUID;size:32;"`  //场馆ID 
    ThumbnailUrl  string `json:"thumbnailUrl" form:"thumbnailUrl" gorm:"column:thumbnail_url;comment:缩略图URL;"`  //缩略图URL 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:场馆名称;size:255;"`  //场馆名称 
    Province  string `json:"province" form:"province" gorm:"column:province;comment:省;size:255;"`  //省 
    City  string `json:"city" form:"city" gorm:"column:city;comment:市;size:255;"`  //市 
    District  string `json:"district" form:"district" gorm:"column:district;comment:区/县;size:255;"`  //区/县 
    Address  string `json:"address" form:"address" gorm:"column:address;comment:详细地址;"`  //详细地址 
    Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:联系电话;size:20;"`  //联系电话 
}


// TableName 场馆 Venue自定义表名 venue
func (Venue) TableName() string {
    return "venue"
}

// BeforeCreate 创建前
func (venue *Venue) BeforeCreate() (err error) {
	venue.Id = strings.ReplaceAll(uuid.NewString(), "-", "")
	return
}
