package leep

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep"
	leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
	"github.com/google/uuid"
)

type VideosService struct{}

// CreateVideos 创建视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) CreateVideos(userId, videoId string, file *multipart.FileHeader) (videoUrl string, err error) {
	// Generate UUID for new videos only
	if videoId == "" {
		videoId = uuid.NewString()
	}

	filePath := fmt.Sprintf("./temp/%s/%s%s", userId, videoId, path.Ext(file.Filename))

	// Create directory if it doesn't exist
	err = os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("目录创建失败: %s", err.Error())
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %w", err)
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	_, err = io.Copy(dst, src) // Use io.Copy to efficiently copy the file
	if err != nil {
		return "", fmt.Errorf("复制文件失败: %w", err)
	}

	var videos leep.Videos
	videos.Id = videoId
	videos.UserId = userId
	videos.VideoUrl = filePath // Store the relative path
	videos.VideoType = "local"
	videos.CreatedAt = &time.Time{}
	videos.UpdatedAt = &time.Time{}

	err = global.MustGetGlobalDBByDBName("leep").Create(&videos).Error
	if err != nil {
		return "", fmt.Errorf("数据库保存失败: %w", err)
	}

	return filePath, nil
}

// DeleteVideos 删除视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) DeleteVideos(id string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&leep.Videos{}, "id = ?", id).Error
	return err
}

// DeleteVideosByIds 批量删除视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) DeleteVideosByIds(ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&[]leep.Videos{}, "id in ?", ids).Error
	return err
}

// UpdateVideos 更新视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) UpdateVideos(videos leep.Videos) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Model(&leep.Videos{}).Where("id = ?", videos.Id).Updates(&videos).Error
	return err
}

// GetVideos 根据id获取视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) GetVideos(id string) (videos leep.Videos, err error) {
	err = global.MustGetGlobalDBByDBName("leep").Where("id = ?", id).First(&videos).Error
	return
}

// GetVideosInfoList 分页获取视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) GetVideosInfoList(info leepReq.VideosSearch) (list []leep.Videos, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("leep").Model(&leep.Videos{})
	var videoss []leep.Videos
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&videoss).Error
	return videoss, total, err
}
func (videosService *VideosService) GetVideosPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
