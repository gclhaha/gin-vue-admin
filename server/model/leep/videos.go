// 自动生成模板Videos
package leep
import (
	"time"
)

// 视频管理 结构体  Videos
type Videos struct {
    Id  string `json:"id" form:"id" gorm:"primarykey;column:id;comment:视频ID，UUID;size:32;"`  //视频ID，UUID 
    UserId  string `json:"userId" form:"userId" gorm:"column:user_id;comment:上传用户的ID;size:32;"`  //上传用户的ID 
    VideoUrl  string `json:"videoUrl" form:"videoUrl" gorm:"column:video_url;comment:视频路径;size:255;"`  //视频路径 
    VideoType  string `json:"videoType" form:"videoType" gorm:"column:video_type;comment:视频类型：local, aliyun, upyun;size:20;"`  //视频类型：local, aliyun, upyun 
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`  //创建时间 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`  //更新时间 
}


// TableName 视频管理 Videos自定义表名 videos
func (Videos) TableName() string {
    return "videos"
}

