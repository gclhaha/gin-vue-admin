// 自动生成模板VideoHighlight
package leep
import (
	"time"
)

// 视频高光 结构体  VideoHighlight
type VideoHighlight struct {
    Id  string `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:32;"`  //id字段 
    VideoId  string `json:"videoId" form:"videoId" gorm:"column:video_id;comment:视频 ID;size:32;"`  //视频 ID 
    UserId  string `json:"userId" form:"userId" gorm:"column:user_id;comment:用户 ID;size:32;"`  //用户 ID 
    HighlightStartTime  *int `json:"highlightStartTime" form:"highlightStartTime" gorm:"column:highlight_start_time;comment:高光片段的开始时间，以秒为单位;size:32;"`  //高光片段的开始时间，以秒为单位 
    HighlightEndTime  *int `json:"highlightEndTime" form:"highlightEndTime" gorm:"column:highlight_end_time;comment:高光片段的结束时间，以秒为单位;size:32;"`  //高光片段的结束时间，以秒为单位 
    HighlightType  string `json:"highlightType" form:"highlightType" gorm:"column:highlight_type;comment:高光类型，用于分类或标识高光的属性;size:50;"`  //高光类型，用于分类或标识高光的属性 
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;size:6;"`  //创建时间 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;size:6;"`  //更新时间 
}


// TableName 视频高光 VideoHighlight自定义表名 video_highlight
func (VideoHighlight) TableName() string {
    return "video_highlight"
}

